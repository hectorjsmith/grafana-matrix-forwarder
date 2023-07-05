package service

import (
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/formatter"
	"grafana-matrix-forwarder/matrix"
	"grafana-matrix-forwarder/model"
	"log"
)

type Forwarder struct {
	AppSettings                cfg.AppSettings
	MatrixWriter               matrix.Writer
	alertToSentEventMap        map[string]sentMatrixEvent
	alertMapPersistenceEnabled bool
}

func NewForwarder(appSettings cfg.AppSettings, writer matrix.Writer) Forwarder {
	return Forwarder{
		AppSettings:                appSettings,
		MatrixWriter:               writer,
		alertToSentEventMap:        map[string]sentMatrixEvent{},
		alertMapPersistenceEnabled: appSettings.PersistAlertMap,
	}
}

func (f *Forwarder) ForwardEvents(roomIds []string, alerts []model.AlertData) error {
	for _, id := range roomIds {
		for _, alert := range alerts {
			err := f.forwardSingleEvent(id, alert)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Forwarder) forwardSingleEvent(roomID string, alert model.AlertData) error {
	log.Printf("alert received (%s) - forwarding to room: %v", alert.Id, roomID)

	resolveWithReaction := f.AppSettings.ResolveMode == cfg.ResolveWithReaction
	resolveWithReply := f.AppSettings.ResolveMode == cfg.ResolveWithReply

	if sentEvent, ok := f.alertToSentEventMap[alert.Id]; ok {
		if alert.State == model.AlertStateResolved && resolveWithReaction {
			return f.sendResolvedReaction(roomID, sentEvent.EventID, alert)
		}
		if alert.State == model.AlertStateResolved && resolveWithReply {
			return f.sendResolvedReply(roomID, sentEvent, alert)
		}
	}
	return f.sendAlertMessage(roomID, alert)
}

func (f *Forwarder) sendResolvedReaction(roomID, eventID string, alert model.AlertData) error {
	reaction := formatter.GenerateReaction(alert)
	f.deleteMatrixEvent(alert.Id)
	_, err := f.MatrixWriter.React(roomID, eventID, reaction)
	return err
}

func (f *Forwarder) sendResolvedReply(roomID string, sentEvent sentMatrixEvent, alert model.AlertData) error {
	rawReply, formattedReply, err := formatter.GenerateReply(sentEvent.SentFormattedBody, alert)
	if err != nil {
		return err
	}
	f.deleteMatrixEvent(alert.Id)
	_, err = f.MatrixWriter.Reply(roomID, sentEvent.EventID, rawReply, formattedReply)
	return err
}

func (f *Forwarder) sendAlertMessage(roomID string, alert model.AlertData) error {
	rawMessage, formattedMessage, err := formatter.GenerateMessage(alert, f.AppSettings.MetricRounding)
	if err != nil {
		return err
	}
	resp, err := f.MatrixWriter.Send(roomID, rawMessage, formattedMessage)
	if err == nil {
		f.storeMatrixEvent(alert.Id, resp, formattedMessage)
	}
	return err
}
