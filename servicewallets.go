package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
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

func (l *LBD) RetrieveBalanceOfAllNonFungiblesServiceWallet(walletAddress, contractId string, pager *Pager) ([]*NonFungible, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/non-fungibles", walletAddress, contractId)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*NonFungible{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

func (l LBD) RetrieveBalanceOfSpecificTypeOfNonFungiblesServiceWallet(walletAddress, contractId, tokenType string, pager *Pager) ([]*NonFungibleToken, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/non-fungibles/%s", walletAddress, contractId, tokenType)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*NonFungibleToken{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	for _, t := range ret {
		t.TokenType = tokenType
	}

	return ret, err
}

// Retrieve

type WalletInformation struct {
	Name          string `json:"name"`
	WalletAddress string `json:"walletAddress"`
	CreatedAt     int64  `json:"createdAt"`
}

func (l *LBD) RetrieveServiceWalletInformation(walletAddress string) (*WalletInformation, error) {
	path := fmt.Sprintf("/v1/wallets/%s", walletAddress)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(WalletInformation)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

func (l *LBD) RetrieveServiceWalletTransactionHistory(walletAddress string) ([]*Transaction, error) {
	path := fmt.Sprintf("/v1/wallets/%s/transactions", walletAddress)

	r := NewGetRequest(path)
	data, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*Transaction{}
	return ret, json.Unmarshal(data.ResponseData, &ret)
}

type RetrieveBaseCoinBalance struct {
	Symbol   string `json:"symbol"`
	Amount   string `json:"amount"`
	Decimals int64  `json:"decimals"`
}

func (l *LBD) RetrieveBaseCoinBalance(walletAddress string) (*RetrieveBaseCoinBalance, error) {
	path := fmt.Sprintf("/v1/wallets/%s/base-coin", walletAddress)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(RetrieveBaseCoinBalance)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

type RetrieveBalanceServiceTokensResponse struct {
	ContractID string `json:"contractId"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Amount     string `json:"amount"`
	Decimals   int64  `json:"decimals"`
	ImgUri     string `json:"imgUri"`
}

func (l *LBD) RetrieveBalanceAllServiceTokens(walletAddress string, pager *Pager) ([]*RetrieveBalanceServiceTokensResponse, error) {
	path := fmt.Sprintf("/v1/wallets/%s/service-tokens", walletAddress)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*RetrieveBalanceServiceTokensResponse{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

func (l *LBD) RetrieveBalanceSpecificServiceTokenWallet(walletAddress, contractId string) (*RetrieveBalanceServiceTokensResponse, error) {
	path := fmt.Sprintf("/v1/wallets/%s/service-tokens/%s", walletAddress, contractId)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(RetrieveBalanceServiceTokensResponse)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

type RetrieveBalanceFungibles struct {
	TokenType string `json:"tokenType"`
	Name      string `json:"name"`
	Meta      string `json:"meta"`
	Amount    string `json:"amount"`
}

func (l *LBD) RetrieveBalanceAllFungibles(walletAddress, contactId string, pager *Pager) ([]*RetrieveBalanceFungibles, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/fungibles", walletAddress, contactId)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*RetrieveBalanceFungibles{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

func (l *LBD) RetrieveBalanceSpecificFungible(walletAddress, contractId, tokenType string) (*RetrieveBalanceFungibles, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/fungibles/%s", walletAddress, contractId, tokenType)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(RetrieveBalanceFungibles)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

type RetrieveBalanceNonFungible struct {
	TokenIndex string `json:"tokenIndex"`
	Name       string `json:"name"`
	Meta       string `json:"meta"`
}

func (l *LBD) RetrieveBalanceSpecificNonFungible(walletAddress, contractId, tokenType, tokenIndex string) (*RetrieveBalanceNonFungible, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/non-fungibles/%s/%s", walletAddress, contractId, tokenType, tokenIndex)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(RetrieveBalanceNonFungible)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

// Transfer

func (r TransferRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?amount=%s&toUserId=%s&walletSecret=%s", base, r.Amount, r.ToUserId, r.WalletSecret)
	}
	return fmt.Sprintf("%s?amount=%s&toAddress=%s&walletSecret=%s", base, r.Amount, r.ToAddress, r.WalletSecret)
}

type TransferRequest struct {
	*Request
	WalletSecret string `json:"walletSecret"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
	Amount       string `json:"amount"`
}

func (l *LBD) TransferServiceTokens(from *Wallet, contractId, to string, amount *big.Int) (*Transaction, error) {
	path := fmt.Sprintf("/v1/wallets/%s/service-tokens/%s/transfer", from.Address, contractId)

	r := TransferRequest{
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

func (l *LBD) TransferFungible(from *Wallet, contractId, to, tokenType string, amount *big.Int) (*Transaction, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/fungibles/%s/transfer", from.Address, contractId, tokenType)

	r := TransferRequest{
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

type TransferList struct {
	TokenId string `json:"tokenId"`
}

type BatchTransferNonFungibleRequest struct {
	*Request
	WalletSecret string          `json:"walletSecret"`
	ToUserId     string          `json:"toUserId,omitempty"`
	ToAddress    string          `json:"toAddress,omitempty"`
	TransferList []*TransferList `json:"transferList"`
}

func (r BatchTransferNonFungibleRequest) Encode() string {
	base := r.Request.Encode()

	tokenIds := make([]string, len(r.TransferList))
	for i, m := range r.TransferList {
		tokenIds[i] = m.TokenId
	}

	transferList := fmt.Sprintf("transferList.tokenId=%s",
		strings.Join(tokenIds, ","),
	)

	if r.ToUserId != "" {
		return fmt.Sprintf("%s?toUserId=%s&%s&walletSecret=%s", base, r.ToUserId, transferList, r.WalletSecret)
	}
	return fmt.Sprintf("%s?toAddress=%s&%s&walletSecret=%s", base, r.ToAddress, transferList, r.WalletSecret)
}

func (l *LBD) BatchTransferNonFungible(walletAddress, contactId, to string, transferList []*TransferList) (*Transaction, error) {
	path := fmt.Sprintf("/v1/wallets/%s/item-tokens/%s/non-fungibles/batch-transfer", walletAddress, contactId)

	r := BatchTransferNonFungibleRequest{
		Request:      NewPostRequest(path),
		WalletSecret: l.Owner.Secret,
		TransferList: transferList,
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
