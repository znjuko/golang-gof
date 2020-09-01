package observer

// Subscriber ...
type Subscriber interface {
	AddEvent(event Event)
}
