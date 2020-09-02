package proxy

import (
	"math/rand"

	"github.com/satori/uuid"
)

// Storage ...
type Storage interface {
	SendRequest(request string) (response string, time int, err error)
}

type storage struct {
	responseStorage map[string]string
	maxTime         int
	minTime         int
}

func (s *storage) SendRequest(request string) (response string, time int, err error) {
	responseGenerator, err := uuid.NewV4()
	if err != nil {
		return "", 0, err
	}

	response = responseGenerator.String()
	time = rand.Intn(s.maxTime-s.minTime) + s.minTime

	s.responseStorage[request] = response

	return
}

// NewStorage ...
func NewStorage(maxTime, minTime int) Storage {
	return &storage{
		responseStorage: make(map[string]string),
		maxTime:         maxTime,
		minTime:         minTime,
	}
}
