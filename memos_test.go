package lbd

import (
	"testing"
)


func TestSaveTheText(t *testing.T) {
	// onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.SaveTheText("てすとだよー",owner)
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveTheText(t *testing.T) {
	is := initializeTest(t)
	txHash := "C1A24B79009E50E8740E4EC697445D62B368F3707074549EA83A0C478E8AA9A3"
	ret, err := l.RetrieveTheText(txHash)
	is.Nil(err)
	t.Log(ret)
}