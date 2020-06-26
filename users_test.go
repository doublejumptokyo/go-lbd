package lbd

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveUserInformation(t *testing.T) {
	ret, err := l.RetrieveUserInformation(userId)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveUserWalletTransactionHistory(t *testing.T) {
	ret, err := l.RetrieveUserWalletTransactionHistory(userId)
	assert.Nil(t, err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestIssueSessionTokenForBaseCoinTransfer(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForBaseCoinTransfer(userId, owner.Address, big.NewInt(1), RequestTypeAOA)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestFuncTransferNonFungibleUserWallet(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.TransferNonFungibleUserWallet(itemTokenContractId, userId, "tlink10ps670a0x6ma5knthjjswgw89d44vmz6xm3umr", tokenType, "00000009")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestIssueSessionTokenForProxySetting(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueSessionTokenForProxySetting(userId, itemTokenContractId, RequestTypeAOA)
	sessionToken = ret.RequestSessionToken
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveSessionTokenStatus(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	ret, err := l.RetrieveSessionTokenStatus(sessionToken)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestCommitTransaction(t *testing.T) {
	if sessionToken == "" {
		t.Skip("Skip because no session token")
	}
	ret, err := l.CommitTransaction(sessionToken)
	assert.Nil(t, err)
	t.Log(ret)
}
