package lbd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveServerTime(t *testing.T) {
	tm, err := l.RetrieveServerTime()
	assert.Nil(t, err)
	t.Logf("Server time: %d", tm)
}
