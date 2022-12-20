package lbd

func (l LBD) RetrieveUserHoldersNonFungibles(userId string, contractId string) ([]*NonFungible, error) {
	all := []*NonFungible{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveBalanceOfAllNonFungiblesUserWallet(userId, contractId, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}

	return all, nil
}

func (l LBD) ListBalanceOfAllServiceTokensUserWallet(userId string) ([]*BalanceOfServiceTokens, error) {
	all := []*BalanceOfServiceTokens{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveBalanceOfAllServiceTokensUserWallet(userId, pager)
		if err != nil {
			return nil, err
		}
		all = append(all, ret...)
		pager.Page++
		if len(ret) < DefaultLimit {
			break
		}
	}

	return all, nil
}

func (l LBD) RetrieveUserHoldersFungibles(userId, contractId string) ([]*BalanceOfFungible, error) {
	all := []*BalanceOfFungible{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveBalanceOfAllFungiblesUserWallet(userId, contractId, pager)
		if err != nil {
			return nil, err
		}
		all = append(all, ret...)
		pager.Page++
		if len(ret) < DefaultLimit {
			break
		}
	}

	return all, nil
}

func (l LBD) ListBalanceOfSpecificTypeOfNonFungiblesUserWallet(userId string, contractId, tokenType string) ([]*NonFungibleToken, error) {
	all := []*NonFungibleToken{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveBalanceOfSpecificTypeOfNonFungiblesUserWallet(userId, contractId, tokenType, pager)
		if err != nil {
			return nil, err
		}
		all = append(all, ret...)
		pager.Page++
		if len(ret) < DefaultLimit {
			break
		}
	}

	return all, nil
}

func (l LBD) RetrieveWalletAddressBalanceOfAllFungibleServiceWallet(walletAddress string, contractId string) ([]*RetrieveBalanceFungibles, error) {
	all := []*RetrieveBalanceFungibles{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveBalanceAllFungibles(walletAddress, contractId, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

func (l LBD) RetrieveWalletAddressBalanceOfAllNonFungiblesServiceWallet(walletAddress string, contractId string) ([]*NonFungible, error) {
	all := []*NonFungible{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveBalanceOfAllNonFungiblesServiceWallet(walletAddress, contractId, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

func (l LBD) RetrieveWalletAddressBalanceOfSpecificTypeOfNonFungiblesServiceWallet(walletAddress, contractId, tokenType string) ([]*NonFungibleToken, error) {
	all := []*NonFungibleToken{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}
	for {
		ret, err := l.RetrieveBalanceOfSpecificTypeOfNonFungiblesServiceWallet(walletAddress, contractId, tokenType, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

// Deprecated: Use ListWalletAddressBalanceAllServiceTokens
func (l LBD) ListWalletAddresseBalanceAllServiceTokens(walletAddress string) ([]*RetrieveBalanceServiceTokensResponse, error) {
	return l.ListWalletAddressBalanceAllServiceTokens(walletAddress)
}

func (l LBD) ListWalletAddressBalanceAllServiceTokens(walletAddress string) ([]*RetrieveBalanceServiceTokensResponse, error) {
	all := []*RetrieveBalanceServiceTokensResponse{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveBalanceAllServiceTokens(walletAddress, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

func (l LBD) ListAllFungiblesItemToken(contractId string) ([]*TokenType, error) {
	all := []*TokenType{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.ListAllFungibles(contractId, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

func (l LBD) RetrieveAllFungibleHoldersItemToken(contractId, tokenType string) ([]*FungibleHolders, error) {
	all := []*FungibleHolders{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveAllFungibleHolders(contractId, tokenType, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

func (l LBD) RetrieveHolderOfSpecificNonFungibleItemToken(contractId, tokenType string) ([]*Holder, error) {
	all := []*Holder{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.RetrieveHolderOfSpecificNonFungible(contractId, tokenType, pager)
		if err != nil {
			return nil, err
		}
		all = append(all, ret...)
		pager.Page++
		if len(ret) < DefaultLimit {
			break
		}
	}
	return all, nil
}

func (l LBD) ListTheChildrenOfNonFungibleItemToken(contractId, tokenType, tokenIndex string) ([]*NonFungibleInformation, error) {
	all := []*NonFungibleInformation{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.ListTheChildrenOfNonFungible(contractId, tokenType, tokenIndex, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

func (l LBD) ListAllNonFungiblesItemToken(contractId string) ([]*TokenType, error) {
	all := []*TokenType{}
	pager := &Pager{
		Page:    1,
		OrderBy: "asc",
		Limit:   DefaultLimit,
	}

	for {
		ret, err := l.ListAllNonFungibles(contractId, pager)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		pager.Page++
	}
	return all, nil
}

func (l LBD) UpdateMultipleNonFungibleTokenIconsCache(tokenIds []string, contractId string) error {
	updateList := make([]*UpdateList, len(tokenIds))
	for i, t := range tokenIds {
		tt, ti := DivideId(t)
		updateList[i] = &UpdateList{
			TokenType:  tt,
			TokenIndex: ti,
		}
	}

	_, err := l.UpdateMultipleNonFungibleTokenIcons(contractId, updateList)
	if err != nil {
		return err
	}

	return nil
}

func (l LBD) UpdateMultipleFungibleTokenIconsCache(tokenIds []string, contractId string) error {
	updateList := make([]*UpdateFungibleList, len(tokenIds))
	for i, t := range tokenIds {
		tt, _ := DivideId(t)
		updateList[i] = &UpdateFungibleList{
			TokenType: tt,
		}

		_, err := l.UpdateMultipleFungibleTokenIcons(contractId, updateList)
		if err != nil {
			return err
		}
	}
	return nil
}

func DivideId(tokenId string) (string, string) {
	return tokenId[0:8], tokenId[8:16]
}
