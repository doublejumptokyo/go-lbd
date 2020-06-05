package lbd

import (
	"os"
	"testing"
)

var (
	itemTokenContractId = os.Getenv("ITEMTOKEN_CONTRACT_ID")
	tokenType           = "10000001"
)

func TestListAllNonFungibles(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.ListAllNonFungibles(itemTokenContractId)
	is.Nil(err)
	t.Log(*ret[1])
}

func TestCreateNonFungible(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	is := initializeTest(t)
	ret, err := l.CreateNonFungible(itemTokenContractId, "NobunagaOda", "Tenkafubu", owner)
	is.Nil(err)
	t.Log(ret)
}

func TestRetrieveNonFungibleInformation(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.RetrieveNonFungibleInformation(itemTokenContractId, tokenType, "00000001")
	is.Nil(err)
	t.Log(ret)
}

func TestMintNonFungible(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	is := initializeTest(t)
	ret, err := l.MintNonFungible(itemTokenContractId, tokenType, "Nobnyaga", "uwawa", toAddress, owner)
	is.Nil(err)
	t.Log(ret)
}

func TestUpdateNonFungibleInformation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	is := initializeTest(t)
	ret, err := l.UpdateNonFungibleInformation(itemTokenContractId, tokenType, "00000001", "aaa", "", owner)
	is.Nil(err)
	t.Log(ret)
}
