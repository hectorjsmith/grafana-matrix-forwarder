package matrix

import (
	"fmt"
	"grafana-matrix-forwarder/grafana"
	"log"
)

// SendAlert sends the provided grafana.AlertPayload to the provided WriteCloser using the provided roomID
func SendAlert(wc WriteCloser, roomID string, alert grafana.AlertPayload) (err error) {
	formattedMessageBody := buildFormattedMessageBodyFromAlert(alert)
	formattedMessage := newSimpleFormattedMessage(formattedMessageBody)
	_, err = wc.Write(roomID, formattedMessage)
	return err
}

func buildFormattedMessageBodyFromAlert(alert grafana.AlertPayload) string {
	var message string
	switch alert.State {
	case grafana.AlertStateAlerting:
		message = buildAlertMessage(alert)
	case grafana.AlertStateResolved:
		message = buildResolvedMessage(alert)
	case grafana.AlertStateNoData:
		message = buildNoDataMessage(alert)
	default:
		log.Printf("alert received with unknown state: %s", alert.State)
		message = buildUnknownStateMessage(alert)
	}
	return message
}

func buildAlertMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("💔 ️<b>ALERT</b><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleURL, alert.RuleName, alert.Message)
}

func buildResolvedMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("💚 ️<b>RESOLVED</b><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleURL, alert.RuleName, alert.Message)
}

func buildNoDataMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("❓️<b>NO DATA</b><ul><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleURL, alert.RuleName, alert.Message)
}

func buildUnknownStateMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("❓️<b>UNKNOWN</b><ul><li>Rule: <a href=\"%s\">%s</a> | %s</li><li>State: <b>%s</b></li></ul>",
		alert.RuleURL, alert.RuleName, alert.Message, alert.State)
}
