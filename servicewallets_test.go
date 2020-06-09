package lbd

import (
	"math/big"
	"testing"
)

func TestListAllServiceWallets(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.ListAllServiceWallets()
	is.Nil(err)

	t.Log(ret[0])
}

func TestTransferBaseCoins(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.TransferBaseCoins(owner, toAddress, big.NewInt(1))
	is.Nil(err)
	t.Log(ret)
}
