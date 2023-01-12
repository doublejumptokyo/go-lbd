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

func TestRetrieveTransactionInformationV2(t *testing.T) {
	h := os.Getenv("TX_HASH")
	if h == "" {
		t.Skip()
	}
	ret, err := l.RetrieveTransactionInformationV2(h)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
