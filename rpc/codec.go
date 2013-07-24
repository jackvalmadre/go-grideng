package rpc

const (
	InputRequest  = "input"
	OutputRequest = "output"
)

// There are two types of request: a request to receive input and request to send output.
// The server does not know the type of the incoming request.
type ClientCodec interface {
	// Write request.
	WriteRequest(typ string, header interface{}, body interface{}) error
	// Read response.
	ReadResponse() (Reader, error)
}

type Reader interface {
	ReadHeader(interface{}) error
	ReadBody(interface{}) error
}

type RequestReader interface {
	Type() string
	Reader
}

// The client does know the type of the incoming response.
type ServerCodec interface {
	// Read request.
	ReadRequest() (RequestReader, error)
	// Write response.
	WriteResponse(header interface{}, body interface{}) error
}

// A request to send output contains an index, and possibly an error (plus the body).
type OutputRequestHeader struct {
	Index int
	Error string
}

// A response to send input contains the index (plus the body).
type InputResponseHeader struct {
	Index int
}
