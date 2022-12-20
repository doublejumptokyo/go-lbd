package lbd

import (
	"math/big"
	"testing"
)

func TestListAllServiceWallets(t *testing.T) {
	ret, err := l.ListAllServiceWallets()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret[0])
}

func TestTransferBaseCoins(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.TransferBaseCoins(owner, toAddress, big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveServiceWalletInformation(t *testing.T) {
	ret, err := l.RetrieveServiceWalletInformation(owner.Address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveServiceWalletTransactionHistory(t *testing.T) {
	ret, err := l.RetrieveServiceWalletTransactionHistory(owner.Address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(ret))
	t.Log(ret)
}

func TestRetrieveBaseCoinBalance(t *testing.T) {
	ret, err := l.RetrieveBaseCoinBalance(owner.Address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceAllServiceTokens(t *testing.T) {
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.RetrieveBalanceAllServiceTokens(owner.Address, pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(ret))
	t.Log(ret)
}

func TestRetrieveBalanceSpecificServiceTokenWallet(t *testing.T) {
	ret, err := l.RetrieveBalanceSpecificServiceTokenWallet(owner.Address, serviceTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceAllFungibles(t *testing.T) {
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.RetrieveBalanceAllFungibles(owner.Address, itemTokenContractId, pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(ret))
	t.Log(ret)
}

func TestRetrieveBalanceSpecificFungible(t *testing.T) {
	ret, err := l.RetrieveBalanceSpecificFungible(owner.Address, itemTokenContractId, fungibleTokenType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceSpecificNonFungible(t *testing.T) {
	ret, err := l.RetrieveBalanceSpecificNonFungible(owner.Address, itemTokenContractId, nonFungibleTokenType, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

// Transfer
func TestTransferServiceTokens(t *testing.T) {
	ret, err := l.TransferServiceTokens(owner, serviceTokenContractId, toAddress, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestTransferFungible(t *testing.T) {
	ret, err := l.TransferFungible(owner, serviceTokenContractId, toAddress, fungibleTokenType, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestBatchTransferNonFungible(t *testing.T) {

	transferList := []*TransferList{
		{
			TokenId: nonFungibleTokenType + "00000001",
		},
		{
			TokenId: nonFungibleTokenType + "00000002",
		},
	}
	ret, err := l.BatchTransferNonFungible(owner.Address, itemTokenContractId, toAddress, transferList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
