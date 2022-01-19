package lbd

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var walletAddress string = "tlink1fr9mpexk5yq3hu6jc0npajfsa0x7tl427fuveq"

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

func TestRetrieveServiceWalletInformation(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveServiceWalletInformation(walletAddress)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveServiceWalletTransactionHistory(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveServiceWalletTransactionHistory(walletAddress)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestRetrieveBaseCoinBalance(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveBaseCoinBalance(walletAddress)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveBalanceAllServiceTokens(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveBalanceAllServiceTokens(walletAddress)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestRetrieveBalanceSpecificServiceTokenWallet(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveBalanceSpecificServiceTokenWallet(walletAddress, serviceTokenContractId)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveBalanceAllFungibles(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveBalanceAllFungibles(walletAddress, itemTokenContractId)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestRetrieveBalanceSpecificFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveBalanceSpecificFungible(walletAddress, itemTokenContractId, tokenType)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveBalanceSpecificNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveBalanceSpecificNonFungible(walletAddress, itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}
