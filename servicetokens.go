package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type ServiceToken struct {
	ContractID   string `json:"contractId"`
	OwnerAddress string `json:"ownerAddress"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	ImgURI       string `json:"imgUri"`
	Meta         string `json:"meta"`
	Decimals     int64  `json:"decimals"`
	CreatedAt    int64  `json:"createdAt"`
	TotalSupply  string `json:"totalSupply"`
	TotalMint    string `json:"totalMint"`
	TotalBurn    string `json:"totalBurn"`
	ServiceID    string `json:"serviceId"`
}

func (l *LBD) ListAllServiceTokens() ([]*ServiceToken, error) {
	path := "/v1/service-tokens"
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := []*ServiceToken{}
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

func (l *LBD) RetrieveServiceTokenInformation(contractId string) (*ServiceToken, error) {
	path := fmt.Sprintf("/v1/service-tokens/%s", contractId)
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := &ServiceToken{}
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

type MintServiceTokenRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Amount       string `json:"amount"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
}

func (r MintServiceTokenRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?amount=%s&ownerAddress=%s&ownerSecret=%s&toUserId=%s", base, r.Amount, r.OwnerAddress, r.OwnerSecret, r.ToUserId)
	}
	return fmt.Sprintf("%s?amount=%s&ownerAddress=%s&ownerSecret=%s&toAddress=%s", base, r.Amount, r.OwnerAddress, r.OwnerSecret, r.ToAddress)
}

func (l *LBD) MintServiceToken(contractId string, to string, amount *big.Int) (*Transaction, error) {
	path := fmt.Sprintf("/v1/service-tokens/%s/mint", contractId)
	r := &MintServiceTokenRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
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
