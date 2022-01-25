package lbd

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllFungibles(t *testing.T) {
	ret, err := l.ListAllFungibles(itemTokenContractId)
	assert.Nil(t, err)
	t.Log(*ret[0])
}

func TestRetrieveFungibleInformation(t *testing.T) {
	ret, err := l.RetrieveFungibleInformation(itemTokenContractId, fungibleTokenType)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveAllFungibleHolders(t *testing.T) {
	ret, err := l.RetrieveAllFungibleHolders(itemTokenContractId, fungibleTokenType)
	assert.Nil(t, err)
	t.Log(*ret[0])
}

func TestListAllNonFungibles(t *testing.T) {
	ret, err := l.ListAllNonFungibles(itemTokenContractId)
	assert.Nil(t, err)
	t.Log(*ret[1])
}

func TestCreateNonFungible(t *testing.T) {
	ret, err := l.CreateNonFungible(itemTokenContractId, "NobunagaOda", "Tenkafubu")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveNonFungibleInformation(t *testing.T) {
	ret, err := l.RetrieveNonFungibleInformation(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestMintNonFungible(t *testing.T) {
	ret, err := l.MintNonFungible(itemTokenContractId, tokenType, "Nobnyaga", "uwawa", toAddress)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestUpdateNonFungibleInformation(t *testing.T) {
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
	ret, err := l.RetrieveTheHolderOfSpecificNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestListTheChildrenOfNonFungible(t *testing.T) {
	ret, err := l.ListTheChildrenOfNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(*ret[0])
}

func TestRetrieveTheParentOfNonFungible(t *testing.T) {
	ret, err := l.RetrieveTheParentOfNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}
func TestRetrieveTheRootOfNonFungible(t *testing.T) {
	ret, err := l.RetrieveTheRootOfNonFungible(itemTokenContractId, tokenType, "00000001")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestRetrieveTheStatusOfMultipleFungibleTokenIcons(t *testing.T) {
	ret, err := l.RetrieveTheStatusOfMultipleFungibleTokenIcons(itemTokenContractId, "63f34026-ffef-4dcf-a512-746f3e512378")
	assert.Nil(t, err)
	t.Log(ret)
}
func TestRetrieveTheStatusOfMultipleNonFungibleTokenIcons(t *testing.T) {
	ret, err := l.RetrieveTheStatusOfMultipleNonFungibleTokenIcons(itemTokenContractId, "df6629c7-f1b6-4a82-81b4-6cc3083c2785")
	assert.Nil(t, err)
	t.Log(ret)
}

func TestUpdateMultipleNonFungibleTokenIcons(t *testing.T) {
	updateList := []*UpdateList{
		{
			TokenType:  "10000001",
			TokenIndex: "10000003",
		},
	}
	ret, err := l.UpdateMultipleNonFungibleTokenIcons(itemTokenContractId, updateList)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestUpdateMultipleFungibleTokenIcons(t *testing.T) {
	updateList := []*UpdateFungibleList{
		{
			TokenType: "10000002",
		},
	}
	ret, err := l.UpdateMultipleFungibleTokenIcons(itemTokenContractId, updateList)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestUpdateFungibleInformation(t *testing.T) {
	ret, err := l.UpdateFungibleInformation(itemTokenContractId, fungibleTokenType, name, meta)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestAttachNonFungibleAnother(t *testing.T) {
	ret, err := l.AttachNonFungibleAnother(itemTokenContractId, tokenType, "00000001", " 1000000600000001", userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestDetachNonFungibleParent(t *testing.T) {
	ret, err := l.DetachNonFungibleParent(itemTokenContractId, tokenType, "00000001", userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestIssueFungible(t *testing.T) {
	ret, err := l.IssueFungible(itemTokenContractId, name, meta)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestMintFungible(t *testing.T) {
	ret, err := l.MintFungible(itemTokenContractId, fungibleTokenType, toAddress, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestBurnFungible(t *testing.T) {
	ret, err := l.BurnFungible(itemTokenContractId, tokenType, toAddress, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func MintMultipleNonFungibleRecipients(t *testing.T) {
	mintList := []*MultiMintList{
		{
			TokenType: "10000006",
			Name:      "testToken",
			Meta:      "test test test",
			ToAddress: toAddress,
			ToUserId:  "",
		},
	}

	ret, err := l.MintMultipleNonFungibleRecipients(itemTokenContractId, mintList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestMintMultipleNonFungible(t *testing.T) {
	mintList := []*MintList{
		{
			TokenType: tokenType,
			Name:      name,
			Meta:      meta,
		},
	}

	ret, err := l.MintMultipleNonFungible(itemTokenContractId, toAddress, mintList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestBurnNonFungible(t *testing.T) {
	ret, err := l.BurnNonFungible(itemTokenContractId, tokenType, "00000001", toAddress)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
