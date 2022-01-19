package lbd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllFungibles(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.ListAllFungibles(itemTokenContractId)
	assert.Nil(t, err)
	t.Log(*ret[1])
}

func TestRetrieveFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveFungibleInformation(itemTokenContractId, tokenType)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveAllFungibleHolders(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveAllFungibleHolders(itemTokenContractId, tokenType)
	assert.Nil(t, err)
	t.Log(*ret[1])
}

func TestListAllNonFungibles(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.ListAllNonFungibles(itemTokenContractId)
	assert.Nil(t, err)
	t.Log(*ret[1])
}

func TestCreateNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.CreateNonFungible(itemTokenContractId, "NobunagaOda", "Tenkafubu")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveNonFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveNonFungibleInformation(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestMintNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.MintNonFungible(itemTokenContractId, tokenType, "Nobnyaga", "uwawa", toAddress)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestUpdateNonFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	m := NewMeta()

	err := m.Set("Name", "ナポレオン")
	if err != nil {
		t.Fatal(err)
	}
	err = m.Set("Attack", "100")
	if err != nil {
		t.Fatal(err)
	}
	err = m.Set("干支", "🐍")
	if err != nil {
		t.Fatal(err)
	}
	meta := m.String()

	ret, err := l.UpdateNonFungibleInformation(itemTokenContractId, "10000002", "00000002", "ナポレオン", meta)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ret)
}

func TestRetrieveTheHolderOfSpecificNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveTheHolderOfSpecificNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestListTheChildrenOfNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.ListTheChildrenOfNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(*ret[1])
}

func TestRetrieveTheParentOfNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveTheParentOfNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}
func TestRetrieveTheRootOfNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveTheRootOfNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveTheStatusOfMultipleFungibleTokenIcons(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveTheStatusOfMultipleFungibleTokenIcons(itemTokenContractId, "101")
	assert.Nil(t, err)
	t.Log(ret)
}
func TestRetrieveTheStatusOfMultipleNonFungibleTokenIcons(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveTheStatusOfMultipleNonFungibleTokenIcons(itemTokenContractId, "101")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestUpdateMultipleFungibleTokenIcons(t *testing.T) {
	onlyTxMode(t)
	updateList := []*UpdateList{
		{
			TokenType:  "10000002",
			TokenIndex: "000004c7",
		},
	}
	ret, err := l.UpdateMultipleFungibleTokenIcons(itemTokenContractId, updateList)
	assert.Nil(t, err)
	t.Log(ret)
}
