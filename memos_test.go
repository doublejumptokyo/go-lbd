package lbd

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	memoMsg    = "てすとだよー"
	memoTxHash = os.Getenv("MEMOS_TX_HASH")
)

func TestSaveText(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.SaveText(memoMsg)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveText(t *testing.T) {
	if memoTxHash == "" {
		t.Skip()
	}
	time.Sleep(2 * time.Second)
	ret, err := l.RetrieveText(memoTxHash)
	assert.Equal(t, memoMsg, ret.Memo)

	assert.Nil(t, err)
	t.Log(ret)
}
