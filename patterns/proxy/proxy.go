package proxy

import "math/rand"

type proxy struct {
	storage        Storage
	cachedRequests map[string]string
	maxTime        int
	minTime        int
}

func (p *proxy) SendRequest(request string) (response string, time int, err error) {
	if storedResponse, ok := p.cachedRequests[request]; ok {
		time = rand.Intn(p.maxTime-p.minTime) + p.minTime

		return storedResponse, time, nil
	}

	response, time, err = p.storage.SendRequest(request)
	p.cachedRequests[request] = response
	return
}

// NewProxy ...
func NewProxy(
	storage Storage,
	maxTime, minTime int,
) Storage {
	return &proxy{
		storage:        storage,
		cachedRequests: make(map[string]string),
		maxTime:        maxTime,
		minTime:        minTime,
	}
}
