package lbd

import (
	"math/big"
	"testing"
)

func TestRetrieveUserInformation(t *testing.T) {
	ret, err := l.RetrieveUserInformation(userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveUserWalletTransactionHistory(t *testing.T) {
	ret, err := l.RetrieveUserWalletTransactionHistory(userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestIssueSessionTokenForBaseCoinTransfer(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForBaseCoinTransfer(userId, owner.Address, big.NewInt(1), RequestTypeAOA)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestFuncTransferNonFungibleUserWallet(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.TransferNonFungibleUserWallet(itemTokenContractId, userId, "tlink10ps670a0x6ma5knthjjswgw89d44vmz6xm3umr", tokenType, "00000009")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestIssueSessionTokenForProxySetting(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForProxySetting(userId, itemTokenContractId, RequestTypeAOA)
	if err != nil {
		t.Fatal(err)
	}
	sessionToken = ret.RequestSessionToken
	t.Log(ret)
}

func TestRetrieveSessionTokenStatus(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	ret, err := l.RetrieveSessionTokenStatus(sessionToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestCommitTransaction(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	ret, err := l.CommitTransaction(sessionToken)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBaseCoinBalanceUserWallet(t *testing.T) {
	ret, err := l.RetrieveBaseCoinBalanceUserWallet(userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfAllServiceTokensUserWallet(t *testing.T) {
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.RetrieveBalanceOfAllServiceTokensUserWallet(userId, pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfSpecificServiceTokenUserWallet(t *testing.T) {
	ret, err := l.RetrieveBalanceOfSpecificServiceTokenUserWallet(userId, serviceTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfAllFungiblesUserWallet(t *testing.T) {
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.RetrieveBalanceOfAllFungiblesUserWallet(userId, itemTokenContractId, pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfSpecificFungibleUserWallet(t *testing.T) {
	ret, err := l.RetrieveBalanceOfSpecificFungibleUserWallet(userId, itemTokenContractId, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfNonFungiblesWithTokenTypeUserWallet(t *testing.T) {
	ret, err := l.RetrieveBalanceOfNonFungiblesWithTokenTypeUserWallet(userId, itemTokenContractId, "asc", "", 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveBalanceOfSpecificNonFungibleUserWallet(t *testing.T) {
	ret, err := l.RetrieveBalanceOfSpecificNonFungibleUserWallet(userId, itemTokenContractId, tokenType, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveWhetherTheServiceTokenProxySet(t *testing.T) {
	ret, err := l.RetrieveWhetherTheServiceTokenProxySet(userId, serviceTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveWhetherTheItemTokenProxySet(t *testing.T) {
	ret, err := l.RetrieveWhetherTheItemTokenProxySet(userId, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestIssueSessionTokenForServiceTokenTransfer(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForServiceTokenTransfer(userId, serviceTokenContractId, owner.Address, big.NewInt(1), RequestTypeAOA)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestIssueSessionTokenForServiceTokenProxySetting(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForServiceTokenProxySetting(userId, serviceTokenContractId, RequestTypeAOA)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestTransferDelegatedServiceTokenUserWallet(t *testing.T) {
	ret, err := l.TransferDelegatedServiceTokenUserWallet(userId, serviceTokenContractId, owner.Address, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestTransferDelegatedFungibleUserWallet(t *testing.T) {
	ret, err := l.TransferDelegatedFungibleUserWallet(userId, itemTokenContractId, tokenType, owner.Address, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestTransferDelegatedNonFungible(t *testing.T) {
	ret, err := l.TransferDelegatedNonFungible(userId, itemTokenContractId, tokenType, "00000001", owner.Address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestBatchTransferDelegatedNonFungiblesUserWallet(t *testing.T) {
	ret, err := l.BatchTransferDelegatedNonFungiblesUserWallet(userId, itemTokenContractId, owner.Address, []TransferList{{TokenId: tokenType + "00000001"}, {TokenId: tokenType + "00000002"}})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
