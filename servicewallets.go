package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Wallet struct {
	Name          string `json:"name"`
	WalletAddress string `json:"walletAddress"`
	CreatedAt     int64  `json:"createdAt"`
}

func (l LBD) ListAllServiceWallets() ([]*Wallet, error) {
	path := fmt.Sprintf("/v1/wallets")
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

func (l *LBD) TransferBaseCoins(from *Account, to string, amount *big.Int) (*Transaction, error) {
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
