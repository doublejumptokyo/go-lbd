package lbd

import (
	"os"
	"testing"
)

func TestRetrieveTransactionInformation(t *testing.T) {
	if os.Getenv("TX_HASH") == "" {
		t.Skip()
	}
	ret, err := l.RetrieveTransactionInformation(os.Getenv("TX_HASH"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
