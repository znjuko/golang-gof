package observer

const (
	// EmptyPlace ...
	EmptyPlace = -1
)

type Publisher interface {
	AddSubscriber(subs Subscriber)
	DeleteSubscriber(subsID int)
	Notify(event Event)
}

type publisher struct {
	listners []Subscriber
}

func (p *publisher) AddSubscriber(subs Subscriber) {
	p.listners = append(p.listners, subs)
}

func (p *publisher) DeleteSubscriber(subsID int) {
	place := p.getDeletedPlace(subsID)

	if place == EmptyPlace {
		return
	}

	if place == len(p.listners)-1 {
		p.listners = p.listners[:place]
		return
	}
	p.listners = append(p.listners[:place], p.listners[place+1:]...)
}

func (p *publisher) Notify(event Event) {
	for iter := range p.listners {
		p.listners[iter].AddEvent(event)
	}
}

func (p *publisher) getDeletedPlace(subsID int) (place int) {
	place = EmptyPlace

	for iter := range p.listners {
		if p.listners[iter].GetID() == subsID {
			return iter
		}
	}

	return
}

// NewPublisher
func NewPublisher() Publisher {
	return &publisher{}
}
