package matrix

import (
	"maunium.net/go/mautrix"
)

// WriteCloser allows writing JSON data to a matrix room and closing the connection
type WriteCloser interface {
	// Close the matrix connection
	Close() error
	// GetWriter instance to allow writing data to a matrix room
	GetWriter() Writer
}

type Writer interface {
	// Send a message payload to a given room and get back the response data
	Send(roomID string, contentJSON interface{}) (*mautrix.RespSendEvent, error)
	// React to a given message
	React(roomID string, eventID string, reaction string) (*mautrix.RespSendEvent, error)
}
