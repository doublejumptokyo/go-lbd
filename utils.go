package lbd

type UserHoldersNonFungible struct {
	userId       string
	NonFungibles []*NonFungible
}

func (l LBD) RetrieveUserHoldersNonFungibles(userIds []string, contractId, tokenType string) ([]*UserHoldersNonFungible, error) {
	all := []*UserHoldersNonFungible{}
	for _, u := range userIds {
		uret, err := l.RetrieveBalanceOfAllNonFungiblesUserWallet(u, contractId)
		if err != nil {
			return nil, err
		}
		d := &UserHoldersNonFungible{
			userId:       u,
			NonFungibles: uret,
		}
		all = append(all, d)
	}
	return all, nil
}

func (l LBD) ListBalanceOfAllServiceTokensUserWallet(userIds []string) ([]*BalanceOfServiceTokens, error) {
	all := []*BalanceOfServiceTokens{}
	for _, u := range userIds {
		ret, err := l.RetrieveBalanceOfAllServiceTokensUserWallet(u)
		if err != nil {
			continue
		}
		for _, r := range ret {
			b := &BalanceOfServiceTokens{
				ContractId: r.ContractId,
				Name:       r.Name,
				Symbol:     r.Symbol,
				ImgUri:     r.ImgUri,
				Decimals:   r.Decimals,
				Amount:     r.Amount,
			}
			all = append(all, b)
		}
	}
	return all, nil
}

type UserHoldersFungible struct {
	UserId    string
	Fungibles []*BalanceOfFungible
}

func (l LBD) RetrieveUserHoldersFungibles(userIds []string, contractId, tokenType string) ([]*UserHoldersFungible, error) {
	all := []*UserHoldersFungible{}
	for _, u := range userIds {
		uret, err := l.RetrieveBalanceOfAllFungiblesUserWallet(u, contractId)
		if err != nil {
			continue
		}
		d := &UserHoldersFungible{
			UserId:    u,
			Fungibles: uret,
		}
		all = append(all, d)
	}
	return all, nil
}

type UserNonFungibleTokens struct {
	UserId           string
	NonFungibleToken []*NonFungibleToken
}

func (l LBD) ListBalanceOfSpecificTypeOfNonFungiblesUserWallet(userIds []string, contractId, tokenType string) ([]*UserNonFungibleTokens, error) {
	all := []*UserNonFungibleTokens{}
	for _, u := range userIds {
		ret, err := l.RetrieveBalanceOfSpecificTypeOfNonFungiblesUserWallet(u, contractId, tokenType)
		if err != nil {
			continue
		}
		n := &UserNonFungibleTokens{
			UserId:           u,
			NonFungibleToken: ret,
		}
		all = append(all, n)
	}
	return all, nil
}

type UserBalanceOfNonFungiblesTokenType struct {
	UserId                         string
	BalanceOfNonFungiblesTokenType *BalanceOfNonFungiblesTokenType
}

func (l LBD) ListBalanceOfNonFungiblesWithTokenTypeUserWallet(userIds []string, contractId, orderBy, pageToken string, limit uint32) ([]*UserBalanceOfNonFungiblesTokenType, error) {
	all := []*UserBalanceOfNonFungiblesTokenType{}
	for _, u := range userIds {
		ret, err := l.RetrieveBalanceOfNonFungiblesWithTokenTypeUserWallet(u, contractId, orderBy, pageToken, limit)
		if err != nil {
			continue
		}
		b := &UserBalanceOfNonFungiblesTokenType{
			UserId:                         u,
			BalanceOfNonFungiblesTokenType: ret,
		}
		all = append(all, b)
	}
	return all, nil
}

type UserTransactionHistory struct {
	UserId      string
	Transaction []*Transaction
}

func (l LBD) ListUserWalletTransactionHistoryUserWallet(userIds []string) ([]*UserTransactionHistory, error) {
	all := []*UserTransactionHistory{}
	for _, u := range userIds {
		ret, err := l.RetrieveUserWalletTransactionHistory(u)
		if err != nil {
			continue
		}
		t := &UserTransactionHistory{
			UserId:      u,
			Transaction: ret,
		}
		all = append(all, t)
	}
	return all, nil
}

type WalletAddressFungibleTokens struct {
	WalletAddress string
	FungibleToken []*RetrieveBalanceFungibles
}

func (l LBD) RetrieveWalletAddressBalanceOfAllFungibleServiceWallet(wallets []string, contractId string) ([]*WalletAddressFungibleTokens, error) {
	all := []*WalletAddressFungibleTokens{}
	for _, w := range wallets {
		ret, err := l.RetrieveBalanceAllFungibles(w, contractId)
		if err != nil {
			continue
		}
		u := &WalletAddressFungibleTokens{
			WalletAddress: w,
			FungibleToken: ret,
		}
		all = append(all, u)
	}
	return all, nil
}

type WalletAddressNonFungibleTokens struct {
	WalletAddress    string
	NonFungibleToken []*NonFungible
}

func (l LBD) RetrieveWalletAddressBalanceOfAllNonFungiblesServiceWallet(wallets []string, contractId string) ([]*WalletAddressNonFungibleTokens, error) {
	all := []*WalletAddressNonFungibleTokens{}
	for _, w := range wallets {
		ret, err := l.RetrieveBalanceOfAllNonFungiblesServiceWallet(w, contractId)
		if err != nil {
			continue
		}
		u := &WalletAddressNonFungibleTokens{
			WalletAddress:    w,
			NonFungibleToken: ret,
		}
		all = append(all, u)
	}
	return all, nil
}

type WalletAddressTransactionHistory struct {
	WalletAddress string
	Transaction   []*Transaction
}

func (l LBD) ListServiceWalletTransactionHistory(wallets []string) ([]*WalletAddressTransactionHistory, error) {
	all := []*WalletAddressTransactionHistory{}
	for _, w := range wallets {
		ret, err := l.RetrieveServiceWalletTransactionHistory(w)
		if err != nil {
			continue
		}
		t := &WalletAddressTransactionHistory{
			WalletAddress: w,
			Transaction:   ret,
		}
		all = append(all, t)
	}
	return all, nil
}

type WalletAddressServiceTokens struct {
	WalletAddress string
	ServiceTokens []*RetrieveBalanceServiceTokensResponse
}

func (l LBD) ListWalletAddresseBalanceAllServiceTokens(wallets []string) ([]*WalletAddressServiceTokens, error) {
	all := []*WalletAddressServiceTokens{}
	for _, w := range wallets {
		ret, err := l.RetrieveBalanceAllServiceTokens(w)
		if err != nil {
			continue
		}
		s := &WalletAddressServiceTokens{
			WalletAddress: w,
			ServiceTokens: ret,
		}
		all = append(all, s)
	}
	return all, nil
}

func (l LBD) ListNonFungibleTokenType(contractId string, tokenTypes []string) ([]*TokenType, error) {
	all := []*TokenType{}
	for _, t := range tokenTypes {
		ret, err := l.RetrieveNonFungibleTokenType(contractId, t, nil)
		if err != nil {
			continue
		}
		all = append(all, ret)
	}
	return all, nil
}

func (l LBD) ListFungibleInformation(contractId string, tokenTypes []string) ([]*FungibleInformation, error) {
	all := []*FungibleInformation{}
	for _, t := range tokenTypes {
		ret, err := l.RetrieveFungibleInformation(contractId, t)
		if err != nil {
			continue
		}
		all = append(all, ret)
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
