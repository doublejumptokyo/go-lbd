package lbd

import (
	"testing"
)

func TestRetrieveServiceInformation(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.RetrieveServiceInformation(serviceID)
	is.Nil(err)
	t.Log(ret)
}
