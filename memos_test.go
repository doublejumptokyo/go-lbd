package lbd

import (
	"os"
	"testing"
	"time"
)

var (
	memoMsg    = "てすとだよー"
	memoTxHash = os.Getenv("MEMOS_TX_HASH")
)

func TestSaveText(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.SaveText(memoMsg, owner)
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveText(t *testing.T) {
	is := initializeTest(t)
	if memoTxHash == "" {
		t.Skip()
	}
	time.Sleep(2 * time.Second)
	ret, err := l.RetrieveText(memoTxHash)
	is.Equal(memoMsg, ret.Memo)

	is.Nil(err)
	t.Log(ret)
}