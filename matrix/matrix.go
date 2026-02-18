package matrix

import (
	"context"
	"log"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

// NewMatrixWriteCloser logs in to the provided matrix server URL using the provided user ID and password
// and returns a matrix WriteCloser
func NewMatrixWriteCloser(userID, userPassword, homeserverURL string) (WriteCloser, error) {
	client, err := mautrix.NewClient(homeserverURL, id.UserID(userID), "")
	if err != nil {
		return nil, err
	}

	log.Print("logging into matrix with username + password")
	_, err = client.Login(context.Background(), &mautrix.ReqLogin{
		Type: "m.login.password",
		Identifier: mautrix.UserIdentifier{
			Type: "m.id.user",
			User: userID,
		},
		Password:                 userPassword,
		InitialDeviceDisplayName: "",
		StoreCredentials:         true,
	})
	return buildMatrixWriteCloser(client, true), err
}

// NewMatrixWriteCloser creates a new WriteCloser with the provided user ID and token
func NewMatrixWriteCloserWithToken(userID, token, homeserverURL string) (WriteCloser, error) {
	log.Print("using matrix auth token")
	client, err := mautrix.NewClient(homeserverURL, id.UserID(userID), token)
	if err != nil {
		return nil, err
	}
	return buildMatrixWriteCloser(client, false), err
}

// buildMatrixWriteCloser builds a WriteCloser from a raw matrix client
func buildMatrixWriteCloser(matrixClient *mautrix.Client, closeable bool) WriteCloser {
	return writeCloser{
		writer: writer{
			matrixClient: matrixClient,
		},
		closeable: closeable,
	}
}

type writeCloser struct {
	writer    writer
	closeable bool
}

type writer struct {
	matrixClient *mautrix.Client
}

func (wc writeCloser) GetWriter() Writer {
	return wc.writer
}

func (wc writeCloser) Close() error {
	if !wc.closeable {
		return nil
	}
	_, err := wc.writer.matrixClient.Logout(context.Background())
	return err
}

func buildFormattedMessagePayload(body FormattedMessage) *event.MessageEventContent {
	return &event.MessageEventContent{
		MsgType:       "m.text",
		Body:          body.TextBody,
		Format:        "org.matrix.custom.html",
		FormattedBody: body.HtmlBody,
	}
}

func (w writer) Send(roomID string, body FormattedMessage) (string, error) {
	payload := buildFormattedMessagePayload(body)
	resp, err := w.sendPayload(roomID, event.EventMessage, payload)
	if err != nil {
		return "", err
	}
	return resp.EventID.String(), err
}

func (w writer) Reply(roomID string, eventID string, body FormattedMessage) (string, error) {
	payload := buildFormattedMessagePayload(body)
	payload.RelatesTo = &event.RelatesTo{
		EventID: id.EventID(eventID),
		Type:    event.RelReference,
	}
	resp, err := w.sendPayload(roomID, event.EventMessage, &payload)
	if err != nil {
		return "", err
	}
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
	if err != nil {
		return "", err
	}
	return resp.EventID.String(), err
}

func (w writer) sendPayload(roomID string, eventType event.Type, messagePayload interface{}) (*mautrix.RespSendEvent, error) {
	return w.matrixClient.SendMessageEvent(context.Background(), id.RoomID(roomID), eventType, messagePayload)
}
