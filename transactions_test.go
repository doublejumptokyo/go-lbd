package lbd

import (
	"os"
	"testing"
)

func TestRetrieveTransactionInformation(t *testing.T) {
	is := initializeTest(t)
	if os.Getenv("TX_HASH") == "" {
		t.Skip()
	}
	ret, err := l.RetrieveTransactionInformation(os.Getenv("TX_HASH"))
	is.Nil(err)
	t.Log(ret)
}
