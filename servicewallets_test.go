package lbd

import (
	"math/big"
	"testing"
)

func TestListAllServiceWallets(t *testing.T) {
	assert := initializeTest(t)
	ret, err := l.ListAllServiceWallets()
	assert.Nil(err)

	t.Log(ret[0])
}

func TestTransferBaseCoins(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.TransferBaseCoins(owner, toAddress, big.NewInt(1))
	assert.Nil(err)
	t.Log(ret)
}
