package matrix

import (
	"log"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

// NewMatrixWriteCloser logs in to the provided matrix server URL using the provided user ID and password
// and returns a matrix WriteCloser
func NewMatrixWriteCloser(userID, userPassword, homeserverURL string) (WriteCloser, error) {
	log.Print("starting matrix client ...")

	client, err := mautrix.NewClient(homeserverURL, id.UserID(userID), "")
	if err != nil {
		return nil, err
	}

	_, err = client.Login(&mautrix.ReqLogin{
		Type: "m.login.password",
		Identifier: mautrix.UserIdentifier{
			Type: "m.id.user",
			User: userID,
		},
		Password:                 userPassword,
		InitialDeviceDisplayName: "",
		StoreCredentials:         true,
	})
	return buildMatrixWriteCloser(client), err
}

// buildMatrixWriteCloser builds a WriteCloser from a raw matrix client
func buildMatrixWriteCloser(matrixClient *mautrix.Client) WriteCloser {
	return writeCloser{writer: writer{matrixClient: matrixClient}}
}

type writeCloser struct {
	writer writer
}

type writer struct {
	matrixClient *mautrix.Client
}

func (wc writeCloser) GetWriter() Writer {
	return wc.writer
}

func (wc writeCloser) Close() error {
	_, err := wc.writer.matrixClient.Logout()
	return err
}

func buildFormattedMessagePayload(body string, formattedBody string) *event.MessageEventContent {
	return &event.MessageEventContent{
		MsgType:       "m.text",
		Body:          body,
		Format:        "org.matrix.custom.html",
		FormattedBody: formattedBody,
	}
}

func (w writer) Send(roomID string, body string, formattedBody string) (string, error) {
	payload := buildFormattedMessagePayload(body, formattedBody)
	resp, err := w.sendPayload(roomID, event.EventMessage, payload)
	return resp.EventID.String(), err
}

func (w writer) Reply(roomID string, eventID string, body string, formattedBody string) (string, error) {
	payload := buildFormattedMessagePayload(body, formattedBody)
	payload.RelatesTo = &event.RelatesTo{
		EventID: id.EventID(eventID),
		Type:    event.RelReference,
	}
	resp, err := w.sendPayload(roomID, event.EventMessage, &payload)
	return resp.EventID.String(), err
}

func (w writer) React(roomID string, eventID string, reaction string) (string, error) {
	// Temporary fix to support sending reactions. The key is to pass a pointer to the send method.
	// PR that addresses issue and fix: https://github.com/tulir/mautrix-go/pull/21
	// Fixed by: https://github.com/tulir/mautrix-go/commit/617e6c94cc3a2f046434bf262fadd993daf02141
	payload := event.ReactionEventContent{
		RelatesTo: event.RelatesTo{
			EventID: id.EventID(eventID),
			Type:    event.RelAnnotation,
			Key:     reaction,
		},
	}
	resp, err := w.sendPayload(roomID, event.EventReaction, &payload)
	return resp.EventID.String(), err
}

func (w writer) sendPayload(roomID string, eventType event.Type, messagePayload interface{}) (*mautrix.RespSendEvent, error) {
	return w.matrixClient.SendMessageEvent(id.RoomID(roomID), eventType, messagePayload)
}
