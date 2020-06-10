package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type MemosRequest struct {
	*Request
	WalletAddress string `json:"walletAddress"`
	WalletSecret  string `json:"walletSecret"`
	Memo          string `json:"memo"`
}

func (l *LBD) SaveTheText(fromUserId, to string, amount *big.Int, requestType RequestType) (*SessionToken, error) {
	path := fmt.Sprintf("/v1/users/%s/base-coin/request-transfer?requestType=%s", fromUserId, requestType)
	r := &IssueSessionTokenForBaseCoinTransferRequest{
		Request:     NewPostRequest(path),
		Amount:      amount.String(),
		RequestType: requestType,
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
	return UnmarshalSessionToken(resp.ResponseData)
}

type MemoInformation struct {
	Memo string `json:"memo"`
}

func (l LBD) RetrieveTheText(txHash string) (*MemoInformation, error) {
	r := NewGetRequest("/v1/memos/" + txHash)
	data, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UserInformation)
	return ret, json.Unmarshal(data.ResponseData, ret)
}
