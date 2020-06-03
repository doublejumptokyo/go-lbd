package lbd

import (
	"fmt"
	"testing"
)

func TestRetrieveServerTime(t *testing.T) {
	is := initializeTest(t)
	tm, err := l.RetrieveServerTime()
	is.Nil(err)
	fmt.Println("Server time:", tm)
}
