package lbd

import (
	"math/big"
	"testing"
)


func TestSaveTheText(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.SaveTheText("てすとだよー",owner)
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveTheText(t *testing.T) {
	is := initializeTest(t)
	txHash := "E848200D92C1AD9D12B6A5A044090D32E95B13D5A7668D37D5583E5D53A7EC2F"
	ret, err := l.RetrieveTheText(txHash)
	is.Nil(err)
	t.Log(ret)
}