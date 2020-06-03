package lbd

import (
	"fmt"
	"os"
	"testing"

	"github.com/cheekybits/is"
)

var (
	l = &LBD{}
)

func TestRetrieveServerTime(t *testing.T) {
	is := initializeTest(t)
	tm, err := l.RetrieveServerTime()
	is.Nil(err)
	fmt.Println("Server time:", tm)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	l, err = NewLBD(os.Getenv("API_KEY"))
	is.Nil(err)
	return is
}
