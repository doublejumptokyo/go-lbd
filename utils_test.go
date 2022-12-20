package lbd

import "testing"

func TestRetrieveUserHoldersNonFungibles(t *testing.T) {
	ret, err := l.RetrieveUserHoldersNonFungibles(userId, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListBalanceOfAllServiceTokensUserWallet(t *testing.T) {
	ret, err := l.ListBalanceOfAllServiceTokensUserWallet(userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveUserHoldersFungibles(t *testing.T) {
	ret, err := l.RetrieveUserHoldersFungibles(userId, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListBalanceOfSpecificTypeOfNonFungiblesUserWallet(t *testing.T) {
	ret, err := l.ListBalanceOfSpecificTypeOfNonFungiblesUserWallet(userId, itemTokenContractId, nonFungibleTokenType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveWalletAddressBalanceOfAllNonFungiblesServiceWallet(t *testing.T) {
	ret, err := l.RetrieveWalletAddressBalanceOfAllNonFungiblesServiceWallet(owner.Address, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveWalletAddressBalanceOfSpecificTypeOfNonFungiblesServiceWallet(t *testing.T) {
	ret, err := l.RetrieveWalletAddressBalanceOfSpecificTypeOfNonFungiblesServiceWallet(owner.Address, itemTokenContractId, nonFungibleTokenType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveWalletAddressBalanceOfAllFungibleServiceWallet(t *testing.T) {
	ret, err := l.RetrieveWalletAddressBalanceOfAllFungibleServiceWallet(owner.Address, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListWalletAddresseBalanceAllServiceTokens(t *testing.T) {
	ret, err := l.ListWalletAddresseBalanceAllServiceTokens(owner.Address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListAllFungiblesItemToken(t *testing.T) {
	ret, err := l.ListAllFungiblesItemToken(itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveAllFungibleHoldersItemToken(t *testing.T) {
	ret, err := l.RetrieveAllFungibleHoldersItemToken(itemTokenContractId, fungibleTokenType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestRetrieveHolderOfSpecificNonFungibleItemToken(t *testing.T) {
	ret, err := l.RetrieveHolderOfSpecificNonFungibleItemToken(itemTokenContractId, nonFungibleTokenType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListTheChildrenOfNonFungibleItemToken(t *testing.T) {
	ret, err := l.ListTheChildrenOfNonFungibleItemToken(itemTokenContractId, nonFungibleTokenType, "00000001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestListAllNonFungiblesItemToken(t *testing.T) {
	ret, err := l.ListAllNonFungiblesItemToken(itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestUpdateMultipleNonFungibleTokenIconsCache(t *testing.T) {
	err := l.UpdateMultipleNonFungibleTokenIconsCache([]string{nonFungibleTokenType + "00000001"}, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUdpateMultipleFungibleTokenIconsCache(t *testing.T) {
	err := l.UpdateMultipleFungibleTokenIconsCache([]string{fungibleTokenType + "00000000"}, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
}
