package service

import (
	"grafana-matrix-forwarder/cfg"
	"grafana-matrix-forwarder/formatter"
	"grafana-matrix-forwarder/matrix"
	"grafana-matrix-forwarder/model"
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

func (f *Forwarder) ForwardEvent(roomIds []string, data model.Data) error {
	for _, id := range roomIds {
		err := f.forwardSingleEvent(id, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Forwarder) forwardSingleEvent(roomID string, data model.Data) error {
	resolveWithReaction := f.AppSettings.ResolveMode == cfg.ResolveWithReaction
	resolveWithReply := f.AppSettings.ResolveMode == cfg.ResolveWithReply

	if sentEvent, ok := f.alertToSentEventMap[data.Id]; ok {
		if data.State == model.AlertStateResolved && resolveWithReaction {
			return f.sendResolvedReaction(roomID, sentEvent.EventID, data)
		}
		if data.State == model.AlertStateResolved && resolveWithReply {
			return f.sendResolvedReply(roomID, sentEvent.EventID, data)
		}
	}
	return f.sendAlertMessage(roomID, data)
}

func (f *Forwarder) sendResolvedReaction(roomID, eventID string, data model.Data) error {
	reaction := formatter.GenerateReaction(data)
	f.deleteMatrixEvent(data.Id)
	_, err := f.MatrixWriter.React(roomID, eventID, reaction)
	return err
}

func (f *Forwarder) sendResolvedReply(roomID, eventID string, data model.Data) error {
	rawReply, formattedReply, err := formatter.GenerateReply(data)
	if err != nil {
		return err
	}
	f.deleteMatrixEvent(data.Id)
	_, err = f.MatrixWriter.Reply(roomID, eventID, rawReply, formattedReply)
	return err
}

func (f *Forwarder) sendAlertMessage(roomID string, data model.Data) error {
	rawMessage, formattedMessage, err := formatter.GenerateMessage(data, f.AppSettings.MetricRounding)
	resp, err := f.MatrixWriter.Send(roomID, rawMessage, formattedMessage)
	if err == nil {
		f.storeMatrixEvent(data.Id, resp.EventID.String(), formattedMessage)
	}
	return err
}
