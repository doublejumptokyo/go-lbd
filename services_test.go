package lbd

import (
	"fmt"
	"testing"
)

func TestRetrieveServiceInformation(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.RetrieveServiceInformation(serviceID)
	is.Nil(err)

	fmt.Println(ret)
}
