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
	ret, err := l.TransferBaseCoins(owner, toAddress, big.NewInt(1))
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveServiceWalletInformation(t *testing.T) {
	ret, err := l.RetrieveServiceWalletInformation(owner.Address)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveServiceWalletTransactionHistory(t *testing.T) {
	ret, err := l.RetrieveServiceWalletTransactionHistory(owner.Address)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestRetrieveBaseCoinBalance(t *testing.T) {
	ret, err := l.RetrieveBaseCoinBalance(owner.Address)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveBalanceAllServiceTokens(t *testing.T) {
	ret, err := l.RetrieveBalanceAllServiceTokens(owner.Address)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestRetrieveBalanceSpecificServiceTokenWallet(t *testing.T) {
	ret, err := l.RetrieveBalanceSpecificServiceTokenWallet(owner.Address, serviceTokenContractId)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveBalanceAllFungibles(t *testing.T) {
	ret, err := l.RetrieveBalanceAllFungibles(owner.Address, itemTokenContractId)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestRetrieveBalanceSpecificFungible(t *testing.T) {
	ret, err := l.RetrieveBalanceSpecificFungible(owner.Address, itemTokenContractId, tokenType)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveBalanceSpecificNonFungible(t *testing.T) {
	ret, err := l.RetrieveBalanceSpecificNonFungible(owner.Address, itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}

// Transfer
func TestTransferServiceTokens(t *testing.T) {
	ret, err := l.TransferServiceTokens(owner, serviceTokenContractId, toAddress, big.NewInt(1000))
	assert.Nil(t, err)
	t.Log(ret)
}

func TestTransferFungible(t *testing.T) {
	ret, err := l.TransferFungible(owner, serviceTokenContractId, toAddress, tokenType, big.NewInt(1000))
	assert.Nil(t, err)
	t.Log(ret)
}

func TestBatchTransferNonFungible(t *testing.T) {
	tokenIndex := "00000001"

	transferList := []*TransferList{
		{
			TokenId: tokenType + tokenIndex,
		},
		{
			TokenId: tokenType + tokenIndex,
		},
	}
	ret, err := l.BatchTransferNonFungible(owner.Address, itemTokenContractId, toAddress, transferList)
	assert.Nil(t, err)
	t.Log(ret)
}
