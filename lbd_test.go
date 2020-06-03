package lbd

import (
	"os"
	"testing"

	"github.com/cheekybits/is"
)

var (
	l         = &LBD{}
	serviceID = os.Getenv("SERVICE_ID")
)

func TestSign(t *testing.T) {
	is := is.New(t)
	// https://docs-blockchain.line.biz/api-guide/Authentication
	key := "136db0ad-0fe1-456f-96a4-329be3f93036"
	secret := "9256bf8a-2b86-42fe-b3e0-d3079d0141fe"

	l, err := NewLBD(key, secret)
	is.Nil(err)

	r := &Request{
		nonce:     "Bp0IqgXE",
		timestamp: 1581850266351,
		method:    "GET",
		path:      "/v1/wallets",
	}
	sig := l.Sign(r)

	expected := "2LtyRNI16y/5/RdoTB65sfLkO0OSJ4pCuz2+ar0npkRbk1/dqq1fbt1FZo7fueQl1umKWWlBGu/53KD2cptcCA=="
	is.Equal(sig, expected)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	l, err = NewLBD(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	is.Nil(err)
	return is
}
