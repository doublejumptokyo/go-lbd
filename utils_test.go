package lbd

import "testing"

func TestRetrieveUserHoldersNonFungibles(t *testing.T) {
	ret, err := l.RetrieveUserHoldersNonFungibles([]string{userId}, itemTokenContractId, tokenType)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Log(r.userId)
		for _, m := range r.NonFungibles {
			t.Logf("Name: %s, Meta: %s, TokenType: %s, NumberOfIndex: %s", m.Name, m.Meta, m.TokenType, m.NumberOfIndex)
		}
	}
}

func TestListBalanceOfAllServiceTokensUserWallet(t *testing.T) {
	ret, err := l.ListBalanceOfAllServiceTokensUserWallet([]string{userId, "U69b8d49495102d6eb3148f7b37f2ba5a", "U69b8d49495102d6eb3148f7b37f2ba5e"})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("Name: %s, ContractId: %s, Decimals: %d, ImgUri: %s, Amount: %s, Symbol: %s", r.Name, r.ContractId, r.Decimals, r.ImgUri, r.Amount, r.Symbol)
	}
}

func TestRetrieveUserHoldersFungibles(t *testing.T) {
	ret, err := l.RetrieveUserHoldersFungibles([]string{userId}, itemTokenContractId, fungibleTokenType)
	if err != nil {
		t.Fatal(err)
	}
	for _, m := range ret {
		t.Log(m.UserId)
		for _, f := range m.Fungibles {
			t.Logf("Name: %s, Meta: %s, TokenType: %s, Amount: %s",
				f.Name,
				f.Meta,
				f.TokenType,
				f.Amount,
			)
		}
	}
}

func TestListBalanceOfSpecificTypeOfNonFungiblesUserWallet(t *testing.T) {
	ret, err := l.ListBalanceOfSpecificTypeOfNonFungiblesUserWallet([]string{userId}, itemTokenContractId, tokenType)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Log(r.UserId)
		for _, n := range r.NonFungibleToken {
			t.Logf("Name: %s, Meta: %s, TokenType: %s, TokenIndex: %s",
				n.Name,
				n.Meta,
				n.TokenType,
				n.TokenIndex,
			)
		}
	}
}

func TestListBalanceOfNonFungiblesWithTokenTypeUserWallet(t *testing.T) {
	ret, err := l.ListBalanceOfNonFungiblesWithTokenTypeUserWallet([]string{userId}, itemTokenContractId, "asc", "", 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("UserId: %s", r.UserId)
		t.Logf("NextPageToken: %s", r.BalanceOfNonFungiblesTokenType.NextPageToken)
		t.Logf("PrePageToken: %s", r.BalanceOfNonFungiblesTokenType.PrePageToken)
		for _, b := range r.BalanceOfNonFungiblesTokenType.List {
			t.Logf("Token\n Name: %s, TokenId: %s, CreatedAt: %d, BurnedAt: %d",
				b.Token.Name,
				b.Token.TokenId,
				b.Token.CreatedAt,
				b.Token.BurnedAt,
			)
			t.Logf("Type\n Name: %s, Meta: %s TokenType: %s, CreatedAt: %d, TotalBurn: %s, TotalMint: %s, TotalSupply: %s",
				b.Type.Name,
				b.Type.Meta,
				b.Type.TokenType,
				b.Type.CreatedAt,
				b.Type.TotalBurn,
				b.Type.TotalMint,
				b.Type.TotalSupply,
			)
		}
	}
}

func TestListUserWalletTransactionHistoryUserWallet(t *testing.T) {
	ret, err := l.ListUserWalletTransactionHistoryUserWallet([]string{userId})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("UserId: %s", r.UserId)
		printTransactionHistory(t, r.Transaction)
	}
}

func TestRetrieveWalletAddressBalanceOfAllNonFungiblesServiceWallet(t *testing.T) {
	ret, err := l.RetrieveWalletAddressBalanceOfAllNonFungiblesServiceWallet([]string{owner.Address}, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("WalletAddress: %s", r.WalletAddress)
		for _, n := range r.NonFungibleToken {
			t.Logf("NonFungibleTokens\n Name: %s, Meta: %s, TokenType: %s, NumberOfIndex: %s", n.Name, n.Meta, n.TokenType, n.NumberOfIndex)
		}
	}
}

func TestRetrieveWalletAddressBalanceOfAllFungibleServiceWallet(t *testing.T) {
	ret, err := l.RetrieveWalletAddressBalanceOfAllFungibleServiceWallet([]string{owner.Address}, itemTokenContractId)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("WalletAddress: %s", r.WalletAddress)
		for _, f := range r.FungibleToken {
			t.Logf("FungibleTokens\n Name: %s, Meta: %s, TokenType: %s, Amount: %s", f.Name, f.Meta, f.TokenType, f.Amount)
		}
	}
}

func TestListServiceWalletTransactionHistory(t *testing.T) {
	ret, err := l.ListServiceWalletTransactionHistory([]string{owner.Address})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("WalletAddress: %s", r.WalletAddress)
		printTransactionHistory(t, r.Transaction)
	}
}

func TestListWalletAddresseBalanceAllServiceTokens(t *testing.T) {
	ret, err := l.ListWalletAddresseBalanceAllServiceTokens([]string{owner.Address})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("WalletAddress: %s", r.WalletAddress)
		for _, s := range r.ServiceTokens {
			t.Logf("ServiceTokens \n Name: %s, ContractId: %s, Amount: %s, Symbol: %s, Decimals: %d, ImgUri: %s",
				s.Name,
				s.ContractID,
				s.Amount,
				s.Symbol,
				s.Decimals,
				s.ImgUri,
			)
		}
	}
}

func TestListNonFungibleTokenType(t *testing.T) {
	ret, err := l.ListNonFungibleTokenType(itemTokenContractId, []string{tokenType})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
	for _, r := range ret {
		t.Logf("TokenTypes \n Name: %s, Meta, %s, Token, %+v, TokenType: %s, TotalBurn: %s, TotalMint: %s, TotalSupply: %s, CreatedAt: %d",
			r.Name,
			r.Meta,
			r.Token,
			r.TokenType,
			r.TotalBurn,
			r.TotalMint,
			r.TotalSupply,
			r.CreatedAt,
		)
		for _, k := range r.Token {
			t.Logf("Token \n Name: %s, Meta: %s, TokenIndex: %s, CreatedAt: %d, BurnedAt: %d",
				k.Name,
				k.Meta,
				k.TokenIndex,
				k.CreatedAt,
				k.BurnedAt,
			)
		}
	}
}

func TestListFungibleInformation(t *testing.T) {
	ret, err := l.ListFungibleInformation(itemTokenContractId, []string{fungibleTokenType, "00000002"})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range ret {
		t.Logf("Token \n Name: %s, Meta: %s, TokenType: %s, TotalBurn: %s, TotalMint: %s, TotalSupply: %s, CreatedAt: %d",
			r.Name,
			r.Meta,
			r.TokenType,
			r.TotalBurn,
			r.TotalMint,
			r.TotalSupply,
			r.CreatedAt,
		)
	}
}

func printTransactionHistory(t *testing.T, transactions []*Transaction) {
	for _, a := range transactions {
		t.Logf("Transaction\n Height: %d, TxHash: %s, CodeSpace: %s, Code: %d, Index: %d, GasUsed: %d, GasWanted: %d, TimeStamp: %d",
			a.Height,
			a.Txhash,
			a.Codespace,
			a.Code,
			a.Index,
			a.GasUsed,
			a.GasWanted,
			a.Timestamp,
		)

		t.Logf("Tx\n Type: %s, Memo: %s, Msg: %s, Gas: %d, Amount: %+v",
			a.Tx.Type,
			a.Tx.Value.Memo,
			a.Tx.Value.Msg,
			a.Tx.Value.Fee.Gas,
			a.Tx.Value.Fee.Amount,
		)
		for _, ls := range a.Logs {
			t.Logf("Logs\n Log: %s, Success: %t, MsgIndex: %d, Events: %+v",
				ls.Log,
				ls.Success,
				ls.MsgIndex,
				ls.Events,
			)
		}
	}
}

func TestUpdateMultipleNonFungibleTokenIconsCache(t *testing.T) {
	err := l.UpdateMultipleNonFungibleTokenIconsCache([]string{tokenType + "00000001"}, itemTokenContractId)
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
