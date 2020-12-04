package matrix

import (
	"log"
	"maunium.net/go/mautrix"
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

// CreateClient receives a user ID, password, and server URL and returns a matrix client
func CreateClient(userID, userPassword, homeserverURL string) (*mautrix.Client, error) {
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
	return client, err
}

func newSimpleFormattedMessage(formattedBody string) EventFormattedMessage {
	bodyWithoutParagraphs := htmlParagraphRegex.ReplaceAllString(formattedBody, " ")
	plainBody := htmlTagRegex.ReplaceAllString(bodyWithoutParagraphs, "")
	return newFormattedMessage(plainBody, formattedBody)
}

func newFormattedMessage(body, formattedBody string) EventFormattedMessage {
	return EventFormattedMessage{
		MsgType:       "m.text",
		Body:          body,
		Format:        "org.matrix.custom.html",
		FormattedBody: formattedBody,
	}
}
