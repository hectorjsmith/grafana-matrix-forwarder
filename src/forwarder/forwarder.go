package forwarder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/grafana"
	"grafana-matrix-forwarder/matrix"
	htmlTemplate "html/template"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	textTemplate "text/template"
)

type AlertForwarder struct {
	AppSettings         cfg.AppSettings
	Writer              matrix.Writer
	alertToSentEventMap map[string]sentMatrixEvent
	saveAlertMap        bool
}

type sentMatrixEvent struct {
	EventID           string
	SentFormattedBody string
}

type alertMessageData struct {
	MetricRounding int
	StateStr       string
	StateEmoji     string
	Payload        grafana.AlertPayload
}

const (
	alertMessageTemplateStr = `{{ .StateEmoji }} <b>{{ .StateStr }}</b>
{{- with .Payload }}<p>Rule: <a href="{{ .RuleURL }}">{{ .RuleName }}</a> | {{ .Message }}</p>
{{- if gt (len .EvalMatches) 0 }}<ul>{{ range $match := .EvalMatches }}<li><b>{{ .Metric }}</b>: {{ RoundValue .Value $.MetricRounding }}</li>{{ end }}</ul>{{ end }}
{{- if gt (len .Tags) 0 }}<p>Tags:</p><ul>{{ range $tagKey, $tagValue := .Tags }}<li><b>{{ $tagKey }}</b>: {{ $tagValue }}</li>{{ end }}</ul>{{ end }}{{ end }}`
	resolvedReactionStr    = `✅`
	resolveReplyStr        = "<mx-reply><blockquote>{{ . }}</blockquote></mx-reply>💚 ️<b>RESOLVED</b>"
	resolveReplyPlainStr   = `💚 ️RESOLVED`
	sentMatrixEventMapFile = "grafanaToMatrixMap.json"
)

var (
	htmlTagRegex       = regexp.MustCompile(`<.*?>`)
	htmlParagraphRegex = regexp.MustCompile(`</?p>`)

	alertMessageTemplate = htmlTemplate.Must(htmlTemplate.New("alertMessage").Funcs(htmlTemplate.FuncMap{
		"RoundValue": roundMetricValue,
	}).Parse(alertMessageTemplateStr))
	resolveReplyTemplate = textTemplate.Must(textTemplate.New("resolveReply").Parse(resolveReplyStr))
)

func NewForwarder(appSettings cfg.AppSettings, writer matrix.Writer) *AlertForwarder {
	forwarder := &AlertForwarder{
		AppSettings:         appSettings,
		Writer:              writer,
		alertToSentEventMap: map[string]sentMatrixEvent{},
		saveAlertMap:        true,
	}
	forwarder.prePopulateAlertMap()
	return forwarder
}

// ForwardAlert sends the provided grafana.AlertPayload to the provided matrix.Writer using the provided roomID
func (forwarder *AlertForwarder) ForwardAlert(roomID string, alert grafana.AlertPayload) (err error) {
	resolveWithReaction := forwarder.AppSettings.ResolveMode == cfg.ResolveWithReaction
	resolveWithReply := forwarder.AppSettings.ResolveMode == cfg.ResolveWithReply

	alertID := alert.FullRuleID()
	if sentEvent, ok := forwarder.alertToSentEventMap[alertID]; ok {
		if alert.State == grafana.AlertStateResolved && resolveWithReaction {
			delete(forwarder.alertToSentEventMap, alertID)
			return forwarder.sendReaction(roomID, sentEvent.EventID)
		}
		if alert.State == grafana.AlertStateResolved && resolveWithReply {
			delete(forwarder.alertToSentEventMap, alertID)
			return forwarder.sendReply(roomID, sentEvent)
		}
	}
	return forwarder.sendRegularMessage(roomID, alert, alertID)
}

func (forwarder *AlertForwarder) sendReaction(roomID string, eventID string) (err error) {
	_, err = forwarder.Writer.React(roomID, eventID, resolvedReactionStr)
	return
}

func (forwarder *AlertForwarder) sendReply(roomID string, event sentMatrixEvent) (err error) {
	replyMessageBody, err := executeTextTemplate(resolveReplyTemplate, event.SentFormattedBody)
	if err != nil {
		return
	}
	_, err = forwarder.Writer.Reply(roomID, event.EventID, resolveReplyPlainStr, replyMessageBody)
	return
}

func (forwarder *AlertForwarder) sendRegularMessage(roomID string, alert grafana.AlertPayload, alertID string) (err error) {
	formattedMessageBody, err := buildFormattedMessageBodyFromAlert(alert, forwarder.AppSettings)
	if err != nil {
		return
	}
	plainMessageBody := stripHtmlTagsFromString(formattedMessageBody)
	response, err := forwarder.Writer.Send(roomID, plainMessageBody, formattedMessageBody)
	if err == nil {
		forwarder.alertToSentEventMap[alertID] = sentMatrixEvent{
			EventID:           response.EventID.String(),
			SentFormattedBody: formattedMessageBody,
		}
		forwarder.saveForwardMap()
	}
	return
}

func buildFormattedMessageBodyFromAlert(alert grafana.AlertPayload, settings cfg.AppSettings) (message string, err error) {
	var messageData = alertMessageData{
		StateStr:       "UNKNOWN",
		StateEmoji:     "❓",
		MetricRounding: settings.MetricRounding,
		Payload:        alert,
	}
	switch alert.State {
	case grafana.AlertStateAlerting:
		messageData.StateStr = "ALERT"
		messageData.StateEmoji = "💔"
	case grafana.AlertStateResolved:
		messageData.StateStr = "RESOLVED"
		messageData.StateEmoji = "💚"
	case grafana.AlertStateNoData:
		messageData.StateStr = "NO DATA"
		messageData.StateEmoji = "❓"
	default:
		log.Printf("alert received with unknown state: %s", alert.State)
	}
	return executeAlertTemplate(alertMessageTemplate, messageData)
}

// stripHtmlTagsFromString removes all the HTML tags from an input string.
func stripHtmlTagsFromString(input string) string {
	bodyWithoutParagraphs := htmlParagraphRegex.ReplaceAllString(input, " ")
	plainBody := htmlTagRegex.ReplaceAllString(bodyWithoutParagraphs, "")
	return plainBody
}

func roundMetricValue(rawValue float64, metricRounding int) string {
	var format string
	if metricRounding >= 0 {
		format = fmt.Sprintf("%%.%df", metricRounding)
	} else {
		format = "%v"
	}
	return fmt.Sprintf(format, rawValue)
}

func executeAlertTemplate(template *htmlTemplate.Template, data alertMessageData) (string, error) {
	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, data)
	return buffer.String(), err
}

func executeTextTemplate(template *textTemplate.Template, content string) (string, error) {
	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, content)
	return buffer.String(), err
}

func (forwarder *AlertForwarder) prePopulateAlertMap() {
	fileData, err := ioutil.ReadFile(sentMatrixEventMapFile)
	if err == nil {
		err = json.Unmarshal(fileData, &forwarder.alertToSentEventMap)
	}

	if err != nil {
		log.Printf("failed to load forward map - using blank map (%v)", err)
	}
}

func (forwarder *AlertForwarder) saveForwardMap() {
	if !forwarder.saveAlertMap {
		return
	}

	jsonData, err := json.Marshal(forwarder.alertToSentEventMap)
	if err == nil {
		err = ioutil.WriteFile(sentMatrixEventMapFile, jsonData, os.ModePerm)
	}

	if err != nil {
		log.Printf("failed to save forwarded alert map - functionality disabled (%v)", err)
		forwarder.saveAlertMap = false
	}
}
