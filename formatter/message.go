package formatter

import (
	"grafana-matrix-forwarder/model"
	"log"
)

type alertMessageData struct {
	MetricRounding int
	StateStr       string
	StateEmoji     string
	Payload        model.AlertData
}

func GenerateMessage(alert model.AlertData, metricRounding int) (plainMessage string, formattedMessage string, err error) {
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
	formattedMessage, err = executeHtmlTemplate(alertMessageTemplate, messageData)
	plainMessage = formattedMessageToPlainMessage(formattedMessage)
	return
}

func GenerateReply(originalFormattedMessage string, alert model.AlertData) (plainReply string, formattedReply string, err error) {
	if alert.State == model.AlertStateResolved {
		formattedReply, err = executeTextTemplate(resolveReplyTemplate, originalFormattedMessage)
		plainReply = resolveReplyPlainStr
	}
	return
}
