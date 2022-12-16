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
	ret, err := l.SaveText(memoMsg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveText(t *testing.T) {
	if memoTxHash == "" {
		t.Skip()
	}
	time.Sleep(2 * time.Second)
	ret, err := l.RetrieveText(memoTxHash)
	if err != nil {
		t.Fatal(err)
	}
	if memoMsg != ret.Memo {
		t.Fatalf("not equal:\nexpected: %s\nactual: %s\n", memoMsg, ret.Memo)
	}

	t.Log(ret)
}
