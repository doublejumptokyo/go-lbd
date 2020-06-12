package lbd

import (
	"testing"
)

func TestListAllNonFungibles(t *testing.T) {
	assert := initializeTest(t)
	ret, err := l.ListAllNonFungibles(itemTokenContractId)
	assert.Nil(err)
	t.Log(*ret[1])
}

func TestCreateNonFungible(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.CreateNonFungible(itemTokenContractId, "NobunagaOda", "Tenkafubu")
	assert.Nil(err)
	t.Log(ret)
}

func TestRetrieveNonFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.RetrieveNonFungibleInformation(itemTokenContractId, tokenType, "00000001")
	assert.Nil(err)
	t.Log(ret)
}

func TestMintNonFungible(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.MintNonFungible(itemTokenContractId, tokenType, "Nobnyaga", "uwawa", toAddress)
	assert.Nil(err)
	t.Log(ret)
}

func TestUpdateNonFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.UpdateNonFungibleInformation(itemTokenContractId, tokenType, "00000001", "aaa", "")
	assert.Nil(err)
	t.Log(ret)
}
