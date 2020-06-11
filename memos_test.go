package lbd

import (
	"os"
	"testing"
)


func TestSaveText(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.SaveText("てすとだよー",owner)
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveTheText(t *testing.T) {
	is := initializeTest(t)
	if os.Getenv("MEMOS_TX_HASH") == "" {
		t.Skip()
	}
	ret, err := l.RetrieveTheText(os.Getenv("MEMOS_TX_HASH"))
	is.Nil(err)
	t.Log(ret)
}