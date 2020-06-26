package lbd

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllServiceWallets(t *testing.T) {
	ret, err := l.ListAllServiceWallets()
	assert.Nil(t, err)

	t.Log(ret[0])
}

func TestTransferBaseCoins(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.TransferBaseCoins(owner, toAddress, big.NewInt(1))
	assert.Nil(t, err)
	t.Log(ret)
}
