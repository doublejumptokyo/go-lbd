package lbd

import (
	"encoding/json"
	"fmt"
)

type TokenType struct {
	TokenType   string   `json:"tokenType"`
	Name        string   `json:"name"`
	Meta        string   `json:"meta"`
	CreatedAt   int64    `json:"createdAt"`
	TotalSupply string   `json:"totalSupply"`
	TotalMint   string   `json:"totalMint"`
	TotalBurn   string   `json:"totalBurn"`
	Token       []*Token `json:"token"`
}

type Token struct {
	TokenIndex string `json:"tokenIndex"`
	Name       string `json:"name"`
	Meta       string `json:"meta"`
	CreatedAt  int64  `json:"createdAt"`
	BurnedAt   int64  `json:"burnedAt"`
}

func (l LBD) ListAllNonFungibles(contractId string) ([]*TokenType, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles", contractId)
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := []*TokenType{}
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

type CreateNonFungibleRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Name         string `json:"name"`
	Meta         string `json:"meta"`
}

func (r CreateNonFungibleRequest) Encode() string {
	base := r.Request.Encode()
	return fmt.Sprintf("%s?meta=%s&name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Meta, r.Name, r.OwnerAddress, r.OwnerSecret)
}

func (l *LBD) CreateNonFungible(contractId, name, meta string, owner *Wallet) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles", contractId)
	r := CreateNonFungibleRequest{NewPostRequest(path), owner.Address, owner.Secret, name, meta}
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

type NonFungibleInformation struct {
	Name      string      `json:"name"`
	TokenID   string      `json:"tokenId"`
	Meta      string      `json:"meta"`
	CreatedAt int64       `json:"createdAt"`
	BurnedAt  interface{} `json:"burnedAt"`
}

func (l *LBD) RetrieveNonFungibleInformation(contractId, tokenType, tokenIndex string) (*NonFungibleInformation, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s", contractId, tokenType, tokenIndex)
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(NonFungibleInformation)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

type MintNonFungibleRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Name         string `json:"name"`
	Meta         string `json:"meta"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
}

func (r MintNonFungibleRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?meta=%s&name=%s&ownerAddress=%s&ownerSecret=%s&toUserId=%s", base, r.Meta, r.Name, r.OwnerAddress, r.OwnerSecret, r.ToUserId)
	}
	return fmt.Sprintf("%s?meta=%s&name=%s&ownerAddress=%s&ownerSecret=%s&toAddress=%s", base, r.Meta, r.Name, r.OwnerAddress, r.OwnerSecret, r.ToAddress)
}

func (l *LBD) MintNonFungible(contractId, tokenType, name, meta, to string, owner *Wallet) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/mint", contractId, tokenType)

	r := MintNonFungibleRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: owner.Address,
		OwnerSecret:  owner.Secret,
		Name:         name,
		Meta:         meta,
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

type UpdateNonFungibleInformationRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Name         string `json:"name"`
	Meta         string `json:"meta,omitempty"`
}

func (r UpdateNonFungibleInformationRequest) Encode() string {
	base := r.Request.Encode()
	if r.Meta != "" {
		return fmt.Sprintf("%s?meta=%s&name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Meta, r.Name, r.OwnerAddress, r.OwnerSecret)
	}
	return fmt.Sprintf("%s?name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Name, r.OwnerAddress, r.OwnerSecret)
}

func (l *LBD) UpdateNonFungibleInformation(contractId, tokenType, tokenIndex, name, meta string, owner *Wallet) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s", contractId, tokenType, tokenIndex)

	r := UpdateNonFungibleInformationRequest{
		Request:      NewPutRequest(path),
		OwnerAddress: owner.Address,
		OwnerSecret:  owner.Secret,
		Name:         name,
		Meta:         meta,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}
