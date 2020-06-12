package lbd

import (
	"testing"
)

func TestRetrieveServerTime(t *testing.T) {
	assert := initializeTest(t)
	tm, err := l.RetrieveServerTime()
	assert.Nil(err)
	t.Logf("Server time: %d", tm)
}
