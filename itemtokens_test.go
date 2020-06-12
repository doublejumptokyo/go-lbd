package lbd

import (
	"testing"
)

func TestListAllNonFungibles(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.ListAllNonFungibles(itemTokenContractId)
	is.Nil(err)
	t.Log(*ret[1])
}

func TestCreateNonFungible(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.CreateNonFungible(itemTokenContractId, "NobunagaOda", "Tenkafubu")
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveNonFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.RetrieveNonFungibleInformation(itemTokenContractId, tokenType, "00000001")
	is.Nil(err)
	t.Log(ret)
}

func TestMintNonFungible(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.MintNonFungible(itemTokenContractId, tokenType, "Nobnyaga", "uwawa", toAddress)
	is.Nil(err)
	t.Log(ret)
}

func TestUpdateNonFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	is := initializeTest(t)
	ret, err := l.UpdateNonFungibleInformation(itemTokenContractId, tokenType, "00000001", "aaa", "")
	is.Nil(err)
	t.Log(ret)
}
