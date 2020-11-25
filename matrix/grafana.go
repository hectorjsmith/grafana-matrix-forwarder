package matrix

import (
	"fmt"
	"grafana-matrix-forwarder/grafana"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

func SendAlert(client *mautrix.Client, alert grafana.AlertPayload, roomId string) (err error) {
	formattedMessageBody := buildFormattedMessageBodyFromAlert(alert)
	formattedMessage := newSimpleFormattedMessage(formattedMessageBody)
	_, err = client.SendMessageEvent(id.RoomID(roomId), event.EventMessage, formattedMessage)
	return err
}

func buildFormattedMessageBodyFromAlert(alert grafana.AlertPayload) string {
	var message string
	if alert.State == "alerting" {
		message = buildAlertMessage(alert)
	} else if alert.State == "ok" {
		message = buildResolvedMessage(alert)
	} else if alert.State == "no_data" {
		message = buildNoDataMessage(alert)
	} else {
		log.Printf("alert received with unknown state: %s", alert.State)
		message = buildUnknownStateMessage(alert)
	}
	return message
}

func buildAlertMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("💔 ️<b>ALERT</b><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleUrl, alert.RuleName, alert.Message)
}

func buildResolvedMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("💚 ️<b>RESOLVED</b><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleUrl, alert.RuleName, alert.Message)
}

func buildNoDataMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("❓️<b>NO DATA</b><ul><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleUrl, alert.RuleName, alert.Message)
}

func buildUnknownStateMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("❓️<b>UNKNOWN</b><ul><li>Rule: <a href=\"%s\">%s</a> | %s</li><li>State: <b>%s</b></li></ul>",
		alert.RuleUrl, alert.RuleName, alert.Message, alert.State)
}
