package events

type Fetcher interface {
}

type Processor interface {
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
