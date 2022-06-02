package forwarder

import (
	"bytes"
	"fmt"
	"grafana-matrix-forwarder/cfg"
	htmlTemplate "html/template"
	"log"
	"regexp"
	textTemplate "text/template"
)

type alertMessageData struct {
	MetricRounding int
	StateStr       string
	StateEmoji     string
	Payload        Data
}

const (
	alertMessageTemplateStr = `{{ .StateEmoji }} <b>{{ .StateStr }}</b>
{{- with .Payload }}<p>Rule: <a href="{{ .RuleURL }}">{{ .RuleName }}</a> | {{ .Message }}</p>
{{- if gt (len .EvalMatches) 0 }}<ul>{{ range $match := .EvalMatches }}<li><b>{{ .Metric }}</b>: {{ RoundValue .Value $.MetricRounding }}</li>{{ end }}</ul>{{ end }}
{{- if gt (len .Tags) 0 }}<p>Tags:</p><ul>{{ range $tagKey, $tagValue := .Tags }}<li><b>{{ $tagKey }}</b>: {{ $tagValue }}</li>{{ end }}</ul>{{ end }}{{ end }}`
	resolvedReactionStr  = `✅`
	resolveReplyStr      = "<mx-reply><blockquote>{{ . }}</blockquote></mx-reply>💚 ️<b>RESOLVED</b>"
	resolveReplyPlainStr = `💚 ️RESOLVED`
)

var (
	htmlTagRegex       = regexp.MustCompile(`<.*?>`)
	htmlParagraphRegex = regexp.MustCompile(`</?p>`)

	alertMessageTemplate = htmlTemplate.Must(htmlTemplate.New("alertMessage").Funcs(htmlTemplate.FuncMap{
		"RoundValue": roundMetricValue,
	}).Parse(alertMessageTemplateStr))
	resolveReplyTemplate = textTemplate.Must(textTemplate.New("resolveReply").Parse(resolveReplyStr))
)

func buildFormattedMessageBodyFromAlert(alert Data, settings cfg.AppSettings) (message string, err error) {
	var messageData = alertMessageData{
		StateStr:       "UNKNOWN",
		StateEmoji:     "❓",
		MetricRounding: settings.MetricRounding,
		Payload:        alert,
	}
	switch alert.State {
	case AlertStateAlerting:
		messageData.StateStr = "ALERT"
		messageData.StateEmoji = "💔"
	case AlertStateResolved:
		messageData.StateStr = "RESOLVED"
		messageData.StateEmoji = "💚"
	case AlertStateNoData:
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
