package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	storageMaxTime = 200
	storageMinTime = 20

	proxyMaxTime = 10
	proxyMinTime = 1

	firstRst  = "SELECT * FROM data WHERE id = 1 AND value = 'hello' ORDER BY id"
	secondRst = "DELETE * FROM data WHERE value = 'hello'"
	thirdRst  = "UPDATE data SET value = 'saved from delete' WHERE value = 'hello'"
)

func TestStorage_SendRequest(t *testing.T) {
	storage := NewStorage(storageMaxTime, storageMinTime)
	proxy := NewProxy(storage, proxyMaxTime, proxyMinTime)

	fstStrgResp, frstStrgRespTime, err := proxy.SendRequest(firstRst)
	assert.NoError(t, err)
	fstPrxyResp, frstPrxyRespTime, err := proxy.SendRequest(firstRst)
	assert.NoError(t, err)
	assert.Equal(t, fstStrgResp, fstPrxyResp)
	assert.Less(t, frstPrxyRespTime, frstStrgRespTime)

	sndStrgResp, sndStrgRespTime, err := proxy.SendRequest(secondRst)
	assert.NoError(t, err)
	sndPrxyResp, sndPrxyRespTime, err := proxy.SendRequest(secondRst)
	assert.NoError(t, err)
	assert.Equal(t, sndStrgResp, sndPrxyResp)
	assert.Less(t, sndPrxyRespTime, sndStrgRespTime)

	thrdStrgResp, thrdStrgRespTime, err := proxy.SendRequest(thirdRst)
	assert.NoError(t, err)
	thrdPrxyResp, thrdPrxyRespTime, err := proxy.SendRequest(thirdRst)
	assert.NoError(t, err)
	assert.Equal(t, thrdStrgResp, thrdPrxyResp)
	assert.Less(t, thrdPrxyRespTime, thrdStrgRespTime)
}
