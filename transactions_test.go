package lbd

import (
	"os"
	"testing"
)

func TestRetrieveTransactionInformation(t *testing.T) {
	assert := initializeTest(t)
	if os.Getenv("TX_HASH") == "" {
		t.Skip()
	}
	ret, err := l.RetrieveTransactionInformation(os.Getenv("TX_HASH"))
	assert.Nil(err)
	t.Log(ret)
}
