package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Wallet struct {
	Name      string `json:"name"`
	Address   string `json:"walletAddress"`
	Secret    string `json:"-"`
	CreatedAt int64  `json:"createdAt"`
}

func NewWallet(address, secret string) *Wallet {
	return &Wallet{
		Address: address,
		Secret:  secret,
	}
}

func (l LBD) ListAllServiceWallets() ([]*Wallet, error) {
	path := "/v1/wallets"
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*Wallet{}
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

type TransferBaseCoinsRequest struct {
	*Request
	WalletSecret string `json:"walletSecret"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
	Amount       string `json:"amount"`
}

func (r TransferBaseCoinsRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?amount=%s&toUserId=%s&walletSecret=%s", base, r.Amount, r.ToUserId, r.WalletSecret)
	}
	return fmt.Sprintf("%s?amount=%s&toAddress=%s&walletSecret=%s", base, r.Amount, r.ToAddress, r.WalletSecret)

}

func (l *LBD) TransferBaseCoins(from *Wallet, to string, amount *big.Int) (*Transaction, error) {
	path := fmt.Sprintf("/v1/wallets/%s/base-coin/transfer", from.Address)

	r := TransferBaseCoinsRequest{
		Request:      NewPostRequest(path),
		WalletSecret: from.Secret,
		Amount:       amount.String(),
	}

	if l.IsAddress(to) {
		r.ToAddress = to
	} else {
		r.ToUserId = to
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

type TransferNonFungibleServiceWalletRequest struct {
	*Request
	WalletSecret string `json:"walletSecret"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
}

func (r TransferNonFungibleServiceWalletRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?toUserId=%s&walletSecret=%s", base, r.ToUserId, r.WalletSecret)
	}
	return fmt.Sprintf("%s?toAddress=%s&walletSecret=%s", base, r.ToAddress, r.WalletSecret)
}

func (l *LBD) TransferNonFungibleServiceWallet(walletAddress, walletSecret, contractId, to, tokenType, tokenIndex string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/non-fungibles/%s/%s/transfer", walletAddress, contractId, tokenType, tokenIndex)
	r := &TransferNonFungibleServiceWalletRequest{
		Request:      NewPostRequest(path),
		WalletSecret: walletSecret,
	}

	if l.IsAddress(to) {
		r.ToAddress = to
	} else {
		r.ToUserId = to
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

func (l *LBD) RetrieveBalanceOfAllNonFungiblesServiceWallet(walletAddress, contractId string) ([]*NonFungible, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/non-fungibles", walletAddress, contractId)

	all := []*NonFungible{}
	page := 1
	for {
		r := NewGetRequest(path)
		r.pager.Page = page
		r.pager.OrderBy = "asc"
		resp, err := l.Do(r, true)
		if err != nil {
			return nil, err
		}
		ret := []*NonFungible{}
		err = json.Unmarshal(resp.ResponseData, &ret)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}

		all = append(all, ret...)
		page++
	}
	return all, nil
}

func (l LBD) RetrieveBalanceOfSpecificTypeOfNonFungiblesServiceWallet(walletAddress, contractId, tokenType string) ([]*NonFungibleToken, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/non-fungibles/%s", walletAddress, contractId, tokenType)

	all := []*NonFungibleToken{}
	page := 1
	for {
		r := NewGetRequest(path)
		r.pager.Page = page
		r.pager.OrderBy = "asc"
		resp, err := l.Do(r, true)
		if err != nil {
			return nil, err
		}
		ret := []*NonFungibleToken{}
		err = json.Unmarshal(resp.ResponseData, &ret)
		if err != nil {
			return nil, err
		}
		if len(ret) == 0 {
			break
		}
		all = append(all, ret...)
		page++
	}

	for _, t := range all {
		t.TokenType = tokenType
	}
	return all, nil
}
