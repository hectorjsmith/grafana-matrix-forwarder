package matrix

import (
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"regexp"
)

// EventFormattedMessage is the JSON payload required to send a formatted message in matrix
type EventFormattedMessage struct {
	MsgType       string `json:"msgtype"`
	Body          string `json:"body"`
	Format        string `json:"format"`
	FormattedBody string `json:"formatted_body"`
}

var (
	htmlTagRegex       = regexp.MustCompile(`<.*?>`)
	htmlParagraphRegex = regexp.MustCompile(`</?p>`)
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

// NewSimpleFormattedMessage builds a EventFormattedMessage from a single message string.
// the raw body of the message is set by removing all HTML elements from the formatted body.
func NewSimpleFormattedMessage(formattedBody string) EventFormattedMessage {
	bodyWithoutParagraphs := htmlParagraphRegex.ReplaceAllString(formattedBody, " ")
	plainBody := htmlTagRegex.ReplaceAllString(bodyWithoutParagraphs, "")
	return NewFormattedMessage(plainBody, formattedBody)
}

// NewFormattedMessage builds a EventFormattedMessage from a formatted message and a raw body.
// for best compatibility, the raw body should not contain HTML code.
func NewFormattedMessage(body, formattedBody string) EventFormattedMessage {
	return EventFormattedMessage{
		MsgType:       "m.text",
		Body:          body,
		Format:        "org.matrix.custom.html",
		FormattedBody: formattedBody,
	}
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

func (w writer) Send(roomID string, contentJSON interface{}) (*mautrix.RespSendEvent, error) {
	return w.matrixClient.SendMessageEvent(id.RoomID(roomID), event.EventMessage, contentJSON)
}

func (w writer) React(roomID string, eventID string, reaction string) (*mautrix.RespSendEvent, error) {
	return w.matrixClient.SendReaction(id.RoomID(roomID), id.EventID(eventID), reaction)
}
