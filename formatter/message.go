package formatter

import (
	"grafana-matrix-forwarder/matrix"
	"grafana-matrix-forwarder/model"
	"log"
)

type alertMessageData struct {
	MetricRounding int
	StateStr       string
	StateEmoji     string
	Payload        model.AlertData
}

func GenerateMessage(alert model.AlertData, metricRounding int) (matrix.FormattedMessage, error) {
	var messageData = alertMessageData{
		StateStr:       "UNKNOWN",
		StateEmoji:     "❓",
		MetricRounding: metricRounding,
		Payload:        alert,
	}
	switch alert.State {
	case model.AlertStateAlerting:
		messageData.StateStr = "ALERT"
		messageData.StateEmoji = "💔"
	case model.AlertStateResolved:
		messageData.StateStr = "RESOLVED"
		messageData.StateEmoji = "💚"
	case model.AlertStateNoData:
		messageData.StateStr = "NO DATA"
		messageData.StateEmoji = "❓"
	default:
		log.Printf("alert received with unknown state: %s", alert.State)
	}
	formattedMessage, err := executeHtmlTemplate(alertMessageTemplate, messageData)
	if err != nil {
		return matrix.FormattedMessage{}, err
	}
	plainMessage := formattedMessageToPlainMessage(formattedMessage)
	return matrix.FormattedMessage{
		TextBody: plainMessage,
		HtmlBody: formattedMessage,
	}, err
}

func GenerateReply(originalFormattedMessage string, alert model.AlertData) (matrix.FormattedMessage, error) {
	if alert.State == model.AlertStateResolved {
		formattedReply, err := executeTextTemplate(resolveReplyTemplate, originalFormattedMessage)
		if err != nil {
			return matrix.FormattedMessage{}, err
		}
		plainReply := resolveReplyPlainStr
		return matrix.FormattedMessage{
			TextBody: plainReply,
			HtmlBody: formattedReply,
		}, err
	}
	return matrix.FormattedMessage{}, nil
}
