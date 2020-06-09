package lbd

import "testing"

func TestRetrieveUserInformation(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.RetrieveUserInformation("U8430f7829d8a78aba7f5dcf9a0da9d6c")
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveUserWalletTransactionHistory(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.RetrieveUserWalletTransactionHistory("U8430f7829d8a78aba7f5dcf9a0da9d6c")
	is.Nil(err)
	t.Log(len(ret))
	t.Log(*ret[0])
}

func TestFuncTransferNonFungibleUserWallet(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.TransferNonFungibleUserWallet(itemTokenContractId, "U8430f7829d8a78aba7f5dcf9a0da9d6c", "tlink10ps670a0x6ma5knthjjswgw89d44vmz6xm3umr", tokenType, "00000009", owner)
	is.Nil(err)
	t.Log(ret)
}

func TestIssueSessionTokenForProxySetting(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.IssueSessionTokenForProxySetting("U8430f7829d8a78aba7f5dcf9a0da9d6c", itemTokenContractId, RequestTypeRedirectUri, owner)
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveSessionTokenStatus(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.RetrieveSessionTokenStatus("ciKhkPRDwgdADrCmpoCDPm6T678")
	is.Nil(err)
	t.Log(ret)
}

func TestCommitTransaction(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.CommitTransaction("ciKhkPRDwgdADrCmpoCDPm6T678")
	is.Nil(err)
	t.Log(ret)
}
