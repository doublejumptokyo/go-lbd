package lbd

import (
	"testing"
)

func TestRetrieveServiceInformation(t *testing.T) {
	ret, err := l.RetrieveServiceInformation(serviceID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
