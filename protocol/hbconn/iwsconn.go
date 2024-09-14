package hbconn

type IWSConn interface {
	// WriteMessage is a helper method for getting a writer using NextWriter,
	// writing the message and closing the writer.
	WriteMessage(messageType int, data []byte) error

	// ReadMessage is a helper method for getting a reader using NextReader and
	// reading from that reader to a buffer.
	ReadMessage() (messageType int, p []byte, err error)
}
