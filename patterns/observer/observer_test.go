package observer

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

const (
	subscriberCountDeleteSuccessTest = 4
	subscriberNotifySuccessTest      = 4
	subscriberOldLengthAddNewTest    = 4
	deletedSubscriberIndex           = 3

	firstMsg  = "message1"
	secondMsg = "message2"
	thirdMsg  = "message3"
	fourthMsg = "message4"
)

var (
	events = []Event{
		{Message: firstMsg},
		{Message: secondMsg},
	}
	newEvents = []Event{
		{Message: thirdMsg},
		{Message: fourthMsg},
	}
)

func Test_ObserverDeleteSuccess(t *testing.T) {
	publisher := NewPublisher()

	var subs []Sender
	for iter := 0; iter < subscriberCountDeleteSuccessTest; iter++ {
		subs = append(subs, NewSender(rand.Int()))

		publisher.AddSubscriber(subs[iter])
	}

	endChan := make(chan string, 1)

	eg := errgroup.Group{}
	eg.Go(
		func() (err error) {

			for iter := range events {
				publisher.Notify(events[iter])
			}

			endChan <- ""

			return nil
		},
	)
	eg.Go(
		func() (err error) {
			<-endChan

			for iter := range subs {
				assert.Equal(t, events, subs[iter].SendEvents(len(events)))
			}

			return nil
		},
	)

	err := eg.Wait()
	assert.NoError(t, err)

	publisher.DeleteSubscriber(subs[deletedSubscriberIndex].GetID())

	eg.Go(
		func() (err error) {

			for iter := range newEvents {
				publisher.Notify(newEvents[iter])
			}

			endChan <- ""

			return nil
		},
	)
	eg.Go(
		func() (err error) {
			<-endChan

			for iter := range subs {
				if iter != deletedSubscriberIndex {
					assert.Equal(t, newEvents, subs[iter].SendEvents(len(newEvents)))
					continue
				}

				assert.Equal(t, []Event{}, subs[iter].SendEvents(len(newEvents)))
			}

			return nil
		},
	)

	err = eg.Wait()
	assert.NoError(t, err)
}

func Test_ObserverNotifySuccess(t *testing.T) {
	publisher := NewPublisher()

	var subs []Sender
	for iter := 0; iter < subscriberNotifySuccessTest; iter++ {
		subs = append(subs, NewSender(rand.Int()))

		publisher.AddSubscriber(subs[iter])
	}

	endChan := make(chan string, 1)

	eg := errgroup.Group{}

	eg.Go(
		func() (err error) {

			for iter := range events {
				publisher.Notify(events[iter])
			}

			endChan <- ""

			return nil
		},
	)
	eg.Go(
		func() (err error) {
			<-endChan

			for iter := range subs {
				assert.Equal(t, events, subs[iter].SendEvents(len(events)))
			}

			return nil
		},
	)

	err := eg.Wait()
	assert.NoError(t, err)
}

func Test_ObserverAddNewSubscriberSuccess(t *testing.T) {
	publisher := NewPublisher()

	var subs []Sender
	for iter := 0; iter < subscriberOldLengthAddNewTest; iter++ {
		subs = append(subs, NewSender(rand.Int()))

		publisher.AddSubscriber(subs[iter])
	}

	endChan := make(chan string, 1)

	eg := errgroup.Group{}

	eg.Go(
		func() (err error) {

			for iter := range events {
				publisher.Notify(events[iter])
			}

			endChan <- ""

			return nil
		},
	)
	eg.Go(
		func() (err error) {
			<-endChan

			for iter := range subs {
				assert.Equal(t, events, subs[iter].SendEvents(len(events)))
			}

			return nil
		},
	)

	err := eg.Wait()
	assert.NoError(t, err)

	newSender := NewSender(rand.Int())
	subs = append(subs, newSender)
	publisher.AddSubscriber(newSender)

	eg.Go(
		func() (err error) {

			for iter := range newEvents {
				publisher.Notify(newEvents[iter])
			}

			endChan <- ""

			return nil
		},
	)
	eg.Go(
		func() (err error) {
			<-endChan

			for iter := range subs {
				assert.Equal(t, newEvents, subs[iter].SendEvents(len(newEvents)))
			}

			return nil
		},
	)

	err = eg.Wait()
	assert.NoError(t, err)
}
