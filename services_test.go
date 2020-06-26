package lbd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveServiceInformation(t *testing.T) {
	ret, err := l.RetrieveServiceInformation(serviceID)
	assert.Nil(t, err)
	t.Log(ret)
}
