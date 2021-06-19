package forwarder

import (
	"encoding/json"
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/grafana"
	"grafana-matrix-forwarder/matrix"
	"io/ioutil"
	"log"
	"os"
)

const (
	alertMapFileName = "grafanaToMatrixMap.json"
)

type AlertForwarder struct {
	AppSettings                cfg.AppSettings
	Writer                     matrix.Writer
	alertToSentEventMap        map[string]sentMatrixEvent
	alertMapPersistenceEnabled bool
}

type sentMatrixEvent struct {
	EventID           string
	SentFormattedBody string
}

func NewForwarder(appSettings cfg.AppSettings, writer matrix.Writer) *AlertForwarder {
	forwarder := &AlertForwarder{
		AppSettings:                appSettings,
		Writer:                     writer,
		alertToSentEventMap:        map[string]sentMatrixEvent{},
		alertMapPersistenceEnabled: appSettings.PersistAlertMap,
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
		forwarder.persistAlertMap()
	}
	return
}

func (forwarder *AlertForwarder) prePopulateAlertMap() {
	fileData, err := ioutil.ReadFile(alertMapFileName)
	if err == nil {
		err = json.Unmarshal(fileData, &forwarder.alertToSentEventMap)
	}

	if err != nil {
		log.Printf("failed to load alert map - falling back on an empty map (%v)", err)
	}
}

func (forwarder *AlertForwarder) persistAlertMap() {
	if !forwarder.alertMapPersistenceEnabled {
		return
	}

	jsonData, err := json.Marshal(forwarder.alertToSentEventMap)
	if err == nil {
		err = ioutil.WriteFile(alertMapFileName, jsonData, os.ModePerm)
	}

	if err != nil {
		log.Printf("failed to persist alert map - functionality disabled (%v)", err)
		forwarder.alertMapPersistenceEnabled = false
	}
}
