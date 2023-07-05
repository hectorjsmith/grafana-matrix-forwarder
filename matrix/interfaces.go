package matrix

// WriteCloser allows writing JSON data to a matrix room and closing the connection
type WriteCloser interface {
	// Close the matrix connection
	Close() error
	// GetWriter instance to allow writing data to a matrix room
	GetWriter() Writer
}

type Writer interface {
	// Send a message payload to a given room and get back the response data, returns the event ID if successful
	Send(roomID string, body string, formattedBody string) (string, error)
	// Reply to the provided event ID with the provided plain text and formatted body, returns the event ID if successful
	Reply(roomID string, eventID string, body string, formattedBody string) (string, error)
	// React to a given message, returns the event ID if successful
	React(roomID string, eventID string, reaction string) (string, error)
}
