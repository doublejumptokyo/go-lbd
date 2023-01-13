package lbd

import (
	"math/big"
	"testing"
)

func TestListAllFungibles(t *testing.T) {
	onlyTxMode(t)
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.ListAllFungibles(itemTokenContractId, pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveFungibleInformation(itemTokenContractId, fungibleTokenType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveAllFungibleHolders(t *testing.T) {
	onlyTxMode(t)
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.RetrieveAllFungibleHolders(itemTokenContractId, fungibleTokenType, pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListAllNonFungibles(t *testing.T) {
	onlyTxMode(t)
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.ListAllNonFungibles(itemTokenContractId, pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestCreateNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.CreateNonFungible(itemTokenContractId, "NobunagaOda", "Tenkafubu")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveNonFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveNonFungibleInformation(itemTokenContractId, nonFungibleTokenType, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestMintNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.MintNonFungible(itemTokenContractId, nonFungibleTokenType, "Nobnyaga", "uwawa", toAddress)
	if err != nil {
		t.Fatal(err)
	}
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
	ret, err := l.RetrieveTheHolderOfSpecificNonFungible(itemTokenContractId, nonFungibleTokenType, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListTheChildrenOfNonFungible(t *testing.T) {
	onlyTxMode(t)
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	ret, err := l.ListTheChildrenOfNonFungible(itemTokenContractId, nonFungibleTokenType, "00000001", pager)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveTheParentOfNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveTheParentOfNonFungible(itemTokenContractId, nonFungibleTokenType, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveTheRootOfNonFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.RetrieveTheRootOfNonFungible(itemTokenContractId, nonFungibleTokenType, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateMultipleNonFungibleTokenIcons(t *testing.T) {
	onlyTxMode(t)
	updateList := []*UpdateList{
		{
			TokenType:  "10000001",
			TokenIndex: "10000003",
		},
	}
	ret, err := l.UpdateMultipleNonFungibleTokenIcons(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateMultipleFungibleTokenIcons(t *testing.T) {
	onlyTxMode(t)
	updateList := []*UpdateFungibleList{
		{
			TokenType: "10000002",
		},
	}
	ret, err := l.UpdateMultipleFungibleTokenIcons(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateFungibleInformation(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.UpdateFungibleInformation(itemTokenContractId, fungibleTokenType, name, meta)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestAttachNonFungibleAnother(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.AttachNonFungibleAnother(itemTokenContractId, nonFungibleTokenType, "00000001", " 1000000600000001", userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestDetachNonFungibleParent(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.DetachNonFungibleParent(itemTokenContractId, nonFungibleTokenType, "00000001", userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestIssueFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.IssueFungible(itemTokenContractId, name, meta)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestMintFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.MintFungible(itemTokenContractId, fungibleTokenType, toAddress, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestBurnFungible(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.BurnFungible(itemTokenContractId, fungibleTokenType, toAddress, big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestMintMultipleNonFungibleRecipients(t *testing.T) {
	onlyTxMode(t)
	mintList := []*MultiMintList{
		{
			TokenType: "10000006",
			Name:      "testToken",
			Meta:      "test test test",
			ToAddress: toAddress,
		},
		{
			TokenType: "10000005",
			Name:      "testToken",
			ToAddress: toAddress,
		},
		{
			TokenType: "10000004",
			Name:      "testToken",
			Meta:      "test test test",
			ToUserId:  userId,
		},
		{
			TokenType: "10000003",
			Name:      "testToken",
			ToUserId:  userId,
		},
	}

	ret, err := l.MintMultipleNonFungibleRecipients(itemTokenContractId, mintList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestMintMultipleNonFungible(t *testing.T) {
	onlyTxMode(t)
	mintList := []*MintList{
		{
			TokenType: nonFungibleTokenType,
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
	onlyTxMode(t)
	ret, err := l.BurnNonFungible(itemTokenContractId, nonFungibleTokenType, "00000001", toAddress)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveFungibleTokenMediaResourceStatus(t *testing.T) {
	onlyTxMode(t)

	updateList := []*UpdateFungibleList{{TokenType: "00000002"}}
	updateRet, _ := l.UpdateFungibleTokenMediaResources(itemTokenContractId, updateList)

	ret, err := l.RetrieveFungibleTokenMediaResourceStatus(itemTokenContractId, updateRet.RequestId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveFungibleTokenThumbnailStatus(t *testing.T) {
	updateList := []*UpdateFungibleList{{TokenType: "00000002"}}
	updateRet, err := l.UpdateFungibleTokenThumbnails(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	ret, err := l.RetrieveFungibleTokenThumbnailStatus(itemTokenContractId, updateRet.RequestId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveNonFungibleTokenMediaResourceStatus(t *testing.T) {
	updateList := []*UpdateList{{TokenType: "10000001", TokenIndex: "00000001"}}
	updateRet, err := l.UpdateNonFungibleTokenMediaResources(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	ret, err := l.RetrieveNonFungibleTokenMediaResourceStatus(itemTokenContractId, updateRet.RequestId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveNonFungibleTokenThumbnailStatus(t *testing.T) {
	updateList := []*UpdateList{{TokenType: "10000001", TokenIndex: "00000001"}}
	updateRet, err := l.UpdateNonFungibleTokenThumbnails(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	ret, err := l.RetrieveNonFungibleTokenThumbnailStatus(itemTokenContractId, updateRet.RequestId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateFungibleTokenMediaResources(t *testing.T) {
	updateList := []*UpdateFungibleList{{TokenType: fungibleTokenType}}
	ret, err := l.UpdateFungibleTokenMediaResources(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateFungibleTokenThumbnails(t *testing.T) {
	updateList := []*UpdateFungibleList{{TokenType: fungibleTokenType}}
	ret, err := l.UpdateFungibleTokenThumbnails(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateNonFungibleTokenMediaResources(t *testing.T) {
	updateList := []*UpdateList{{TokenType: nonFungibleTokenType, TokenIndex: "00000001"}}
	ret, err := l.UpdateNonFungibleTokenMediaResources(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateNonFungibleTokenThumbnails(t *testing.T) {
	updateList := []*UpdateList{{TokenType: nonFungibleTokenType, TokenIndex: "00000001"}}
	ret, err := l.UpdateNonFungibleTokenThumbnails(itemTokenContractId, updateList)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
