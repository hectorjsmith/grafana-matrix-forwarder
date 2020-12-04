package matrix

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

// WriteCloser allows writing JSON data to a matrix room and closing the connection
type WriteCloser interface {
	// Close the matrix connection
	Close() error

	// Write a message payload to a given room and get back the response data
	Write(roomID string, contentJSON interface{}) (resp *mautrix.RespSendEvent, err error)
}

// BuildMatrixWriteCloser builds a WriteCloser from a raw matrix client
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
