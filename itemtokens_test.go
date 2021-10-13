package lbd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllNonFungibles(t *testing.T) {
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
