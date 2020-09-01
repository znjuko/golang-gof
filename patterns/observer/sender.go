package observer

import "sync"

// Sender ...
type Sender interface {
	SendEvents(eventCount int) (events []Event)
	GetID() (ID int)
	Subscriber
}

type sender struct {
	mutex  sync.Mutex
	ID     int
	events []Event
}

func (s *sender) AddEvent(event Event) {
	s.mutex.Lock()

	s.events = append(s.events, event)

	s.mutex.Unlock()

	return
}

func (s *sender) SendEvents(eventCount int) (events []Event) {
	s.mutex.Lock()

	if eventCount >= len(s.events) {
		eventCount = len(s.events)
	}

	events = s.events[:eventCount]
	s.events = s.events[eventCount:]

	s.mutex.Unlock()

	return events
}

func (s *sender) GetID() (ID int) {
	return s.ID
}

// NewSender ...
func NewSender(ID int) Sender {
	return &sender{ID: ID}
}
