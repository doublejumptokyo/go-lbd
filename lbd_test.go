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
	fmt.Println("aa")
	is := initializeTest(t)
	_, err := l.RetrieveServerTime()
	is.Nil(err)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	l, err = NewLBD(os.Getenv("API_KEY"))
	is.Nil(err)
	return is
}
