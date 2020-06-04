package lbd

import (
	"testing"
)

func TestRetrieveServerTime(t *testing.T) {
	is := initializeTest(t)
	tm, err := l.RetrieveServerTime()
	is.Nil(err)
	t.Logf("Server time: %d", tm)
}
