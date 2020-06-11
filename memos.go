package lbd

import (
	"encoding/json"
	"fmt"
)

type MemosRequest struct {
	*Request
	WalletAddress string `json:"walletAddress"`
	WalletSecret  string `json:"walletSecret"`
	Memo          string `json:"memo"`
}

func (r MemosRequest) Encode() string {
	base := r.Request.Encode()
	return fmt.Sprintf("%s?memo=%s&walletAddress=%s&walletSecret=%s", base, r.Memo, r.WalletAddress, r.WalletSecret)
}

func (l *LBD) SaveText(memo string, owner *Wallet) (*Transaction, error) {
	path := fmt.Sprintf("/v1/memos")
	r := &MemosRequest{
		Request:       NewPostRequest(path),
		WalletAddress: owner.Address,
		WalletSecret:  owner.Secret,
		Memo:          memo,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

type MemoInformation struct {
	Memo string `json:"memo"`
}

func (l LBD) RetrieveText(txHash string) (*MemoInformation, error) {
	r := NewGetRequest("/v1/memos/" + txHash)
	data, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(MemoInformation)
	return ret, json.Unmarshal(data.ResponseData, ret)
}
