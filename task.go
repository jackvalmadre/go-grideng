package grideng

import "io"

type InputList interface {
	Len() int
	At(i int) Input
}

// An input knows its name and how to serialize itself.
type Input interface {
	// Uniquely identifies an input.
	Name() string
	Write(io.Writer) error
}

// Deserializes an input and returns a task which is given it.
type InputReader interface {
	Read(io.Reader) (Task, error)
}

// A task has an input and a method of execution.
type Task interface {
	Input() Input
	Execute() (Output, error)
}

// An output knows how to serialize and de-serialize itself.
type Output interface {
	Write(io.Writer) error
}

type OutputList interface {
	Read(i int, r io.Reader) error
}
