package lbd

import (
	"math/big"
	"testing"
)

func TestRetrieveUserInformation(t *testing.T) {
	assert := initializeTest(t)
	ret, err := l.RetrieveUserInformation(userId)
	assert.Nil(err)
	t.Log(ret)
}

func TestRetrieveUserWalletTransactionHistory(t *testing.T) {
	assert := initializeTest(t)
	ret, err := l.RetrieveUserWalletTransactionHistory(userId)
	assert.Nil(err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestIssueSessionTokenForBaseCoinTransfer(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.IssueSessionTokenForBaseCoinTransfer(userId, owner.Address, big.NewInt(1), RequestTypeAOA)
	assert.Nil(err)
	t.Log(ret)
}

func TestFuncTransferNonFungibleUserWallet(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.TransferNonFungibleUserWallet(itemTokenContractId, userId, "tlink10ps670a0x6ma5knthjjswgw89d44vmz6xm3umr", tokenType, "00000009")
	assert.Nil(err)
	t.Log(ret)
}

func TestIssueSessionTokenForProxySetting(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.IssueSessionTokenForProxySetting(userId, itemTokenContractId, RequestTypeAOA)
	sessionToken = ret.RequestSessionToken
	assert.Nil(err)
	t.Log(ret)
}

func TestRetrieveSessionTokenStatus(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	assert := initializeTest(t)
	ret, err := l.RetrieveSessionTokenStatus(sessionToken)
	assert.Nil(err)
	t.Log(ret)
}

func TestCommitTransaction(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	assert := initializeTest(t)
	ret, err := l.CommitTransaction(sessionToken)
	assert.Nil(err)
	t.Log(ret)
}
