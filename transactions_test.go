package lbd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveTransactionInformation(t *testing.T) {
	if os.Getenv("TX_HASH") == "" {
		t.Skip()
	}
	ret, err := l.RetrieveTransactionInformation(os.Getenv("TX_HASH"))
	assert.Nil(t, err)
	t.Log(ret)
}
