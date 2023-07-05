package matrix

// WriteCloser allows writing JSON data to a matrix room and closing the connection
type WriteCloser interface {
	// Close the matrix connection
	Close() error
	// GetWriter instance to allow writing data to a matrix room
	GetWriter() Writer
}

type Writer interface {
	// Send a message payload to a given room and get back the event ID if successful
	Send(roomID string, body FormattedMessage) (string, error)
	// Reply to the provided event ID with the provided message, returns the event ID if successful
	Reply(roomID string, eventID string, body FormattedMessage) (string, error)
	// React to a given event ID, returns the new event ID if successful
	React(roomID string, eventID string, reaction string) (string, error)
}

type FormattedMessage struct {
	TextBody string
	HtmlBody string
}
