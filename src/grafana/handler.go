package grafana

import (
	"bytes"
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/matrix"
	htmlTemplate "html/template"
	"log"
	"regexp"
	textTemplate "text/template"
)

type sentMatrixEvent struct {
	eventID           string
	sentFormattedBody string
}

const (
	alertMessageStr      = `💔 ️<b>ALERT</b><p>Rule: <a href="{{ .RuleURL }}">{{ .RuleName }}</a> | {{ .Message }}</p>`
	resolvedMessageStr   = `💚 ️<b>RESOLVED</b><p>Rule: <a href="{{ .RuleURL }}">{{ .RuleName }}</a> | {{ .Message }}</p>`
	noDataMessageStr     = `❓️<b>NO DATA</b><ul><p>Rule: <a href="{{ .RuleURL }}">{{ .RuleName }}</a> | {{ .Message }}</p>`
	unknownMessageStr    = `❓️<b>UNKNOWN</b><ul><li>Rule: <a href="{{ .RuleURL }}">{{ .RuleName }}</a> | {{ .Message }}</li><li>State: <b>{{ .State }}</b></li></ul>`
	resolvedReactionStr  = `✅`
	resolveReplyStr      = "<mx-reply><blockquote>{{ . }}</blockquote></mx-reply>💚 ️<b>RESOLVED</b>"
	resolveReplyPlainStr = `💚 ️RESOLVED`
)

var (
	htmlTagRegex       = regexp.MustCompile(`<.*?>`)
	htmlParagraphRegex = regexp.MustCompile(`</?p>`)

	alertMessageTemplate    = htmlTemplate.Must(htmlTemplate.New("alertMessage").Parse(alertMessageStr))
	resolvedMessageTemplate = htmlTemplate.Must(htmlTemplate.New("resolvedMessage").Parse(resolvedMessageStr))
	noDataMessageTemplate   = htmlTemplate.Must(htmlTemplate.New("noDataMessage").Parse(noDataMessageStr))
	unknownMessageTemplate  = htmlTemplate.Must(htmlTemplate.New("unknownMessage").Parse(unknownMessageStr))
	resolveReplyTemplate    = textTemplate.Must(textTemplate.New("resolveReply").Parse(resolveReplyStr))

	alertToSentEventMap = map[string]sentMatrixEvent{}
)

// ForwardAlert sends the provided grafana.AlertPayload to the provided matrix.Writer using the provided roomID
func ForwardAlert(writer matrix.Writer, roomID string, alert AlertPayload, resolveMode cfg.ResolveMode) (err error) {
	resolveWithReaction := resolveMode == cfg.ResolveWithReaction
	resolveWithReply := resolveMode == cfg.ResolveWithReply

	alertID := alert.FullRuleID()
	if sentEvent, ok := alertToSentEventMap[alertID]; ok {
		if alert.State == AlertStateResolved && resolveWithReaction {
			delete(alertToSentEventMap, alertID)
			return sendReaction(writer, roomID, sentEvent.eventID)
		}
		if alert.State == AlertStateResolved && resolveWithReply {
			delete(alertToSentEventMap, alertID)
			return sendReply(writer, roomID, sentEvent)
		}
	}
	return sendRegularMessage(writer, roomID, alert, alertID)
}

func sendReaction(writer matrix.Writer, roomID string, eventID string) (err error) {
	_, err = writer.React(roomID, eventID, resolvedReactionStr)
	return
}

func sendReply(writer matrix.Writer, roomID string, event sentMatrixEvent) (err error) {
	replyMessageBody, err := executeTextTemplate(resolveReplyTemplate, event.sentFormattedBody)
	if err != nil {
		return
	}
	_, err = writer.Reply(roomID, event.eventID, resolveReplyPlainStr, replyMessageBody)
	return
}

func sendRegularMessage(writer matrix.Writer, roomID string, alert AlertPayload, alertID string) (err error) {
	formattedMessageBody, err := buildFormattedMessageBodyFromAlert(alert)
	if err != nil {
		return
	}
	plainMessageBody := stripHtmlTagsFromString(formattedMessageBody)
	response, err := writer.Send(roomID, plainMessageBody, formattedMessageBody)
	if err == nil {
		alertToSentEventMap[alertID] = sentMatrixEvent{
			eventID:           response.EventID.String(),
			sentFormattedBody: formattedMessageBody,
		}
	}
	return
}

func buildFormattedMessageBodyFromAlert(alert AlertPayload) (message string, err error) {
	switch alert.State {
	case AlertStateAlerting:
		message, err = executeAlertTemplate(alertMessageTemplate, alert)
	case AlertStateResolved:
		message, err = executeAlertTemplate(resolvedMessageTemplate, alert)
	case AlertStateNoData:
		message, err = executeAlertTemplate(noDataMessageTemplate, alert)
	default:
		log.Printf("alert received with unknown state: %s", alert.State)
		message, err = executeAlertTemplate(unknownMessageTemplate, alert)
	}
	return message, err
}

// stripHtmlTagsFromString removes all the HTML tags from an input string.
func stripHtmlTagsFromString(input string) string {
	bodyWithoutParagraphs := htmlParagraphRegex.ReplaceAllString(input, " ")
	plainBody := htmlTagRegex.ReplaceAllString(bodyWithoutParagraphs, "")
	return plainBody
}

func executeAlertTemplate(template *htmlTemplate.Template, alert AlertPayload) (string, error) {
	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, alert)
	return buffer.String(), err
}

func executeTextTemplate(template *textTemplate.Template, content string) (string, error) {
	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, content)
	return buffer.String(), err
}
