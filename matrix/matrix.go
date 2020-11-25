package matrix

import (
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
	"regexp"
)

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

func CreateClient(userId, userPassword, homeserverUrl string) (*mautrix.Client, error) {
	log.Print("starting matrix client ...")

	client, err := mautrix.NewClient(homeserverUrl, id.UserID(userId), "")
	if err != nil {
		return nil, err
	}

	_, err = client.Login(&mautrix.ReqLogin{
		Type: "m.login.password",
		Identifier: mautrix.UserIdentifier{
			Type: "m.id.user",
			User: userId,
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
