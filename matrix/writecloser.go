package matrix

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

type WriteCloser interface {
	Close() error
	Write(roomID string, contentJSON interface{}) (resp *mautrix.RespSendEvent, err error)
}

func BuildMatrixWriteCloser(matrixClient *mautrix.Client) WriteCloser {
	return writeCloser{matrixClient: matrixClient}
}

type writeCloser struct {
	matrixClient *mautrix.Client
}

func (wc writeCloser) Close() error {
	_, err := wc.matrixClient.Logout()
	return err
}

func (wc writeCloser) Write(roomID string, contentJSON interface{}) (resp *mautrix.RespSendEvent, err error) {
	return wc.matrixClient.SendMessageEvent(id.RoomID(roomID), event.EventMessage, contentJSON)
}
