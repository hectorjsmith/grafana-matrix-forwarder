package matrix

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

type Messenger interface {
	Logout() error
	SendMessageEvent(roomID string, contentJSON interface{}) (resp *mautrix.RespSendEvent, err error)
}

func BuildMatrixMessenger(matrixClient *mautrix.Client) Messenger {
	return messengerImpl{matrixClient: matrixClient}
}

type messengerImpl struct {
	matrixClient *mautrix.Client
}

func (msg messengerImpl) Logout() error {
	_, err := msg.matrixClient.Logout()
	return err
}

func (msg messengerImpl) SendMessageEvent(roomID string, contentJSON interface{}) (resp *mautrix.RespSendEvent, err error) {
	return msg.matrixClient.SendMessageEvent(id.RoomID(roomID), event.EventMessage, contentJSON)
}
