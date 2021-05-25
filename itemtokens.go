package lbd

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ItemTokenContractInformation struct {
	ContractID   string `json:"contractId"`
	BaseImgURI   string `json:"baseImgUri"`
	OwnerAddress string `json:"ownerAddress"`
	CreatedAt    int64  `json:"createdAt"`
	ServiceID    string `json:"serviceId"`
}

func (l LBD) RetrieveItemTokenContractInformation(contractId string) (*ItemTokenContractInformation, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s", contractId)
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := &ItemTokenContractInformation{}
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

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

	all := []*TokenType{}
	page := 1
	for {
		r := NewGetRequest(path)
		r.pager.Page = page
		r.pager.OrderBy = "asc"
		resp, err := l.Do(r, true)
		if err != nil {
			return nil, err
		}
		ret := []*TokenType{}
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

func (l *LBD) CreateNonFungible(contractId, name, meta string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles", contractId)
	r := CreateNonFungibleRequest{NewPostRequest(path), l.Owner.Address, l.Owner.Secret, name, meta}
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

type UpdateNonFungibleTokenTypeRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Name         string `json:"name"`
	Meta         string `json:"meta"`
}

func (r UpdateNonFungibleTokenTypeRequest) Encode() string {
	base := r.Request.Encode()
	return fmt.Sprintf("%s?meta=%s&name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Meta, r.Name, r.OwnerAddress, r.OwnerSecret)
}

func (l *LBD) UpdateNonFungibleTokenType(contractId, tokenType, name, meta string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s", contractId, tokenType)
	r := CreateNonFungibleRequest{NewPutRequest(path), l.Owner.Address, l.Owner.Secret, name, meta}
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

type NonFungibleTokenType struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Name         string `json:"name"`
	Meta         string `json:"meta"`
}

func (l *LBD) RetrieveNonFungibleTokenType(contractId, tokenType string) ([]byte, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s", contractId, tokenType)
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return resp.ResponseData, nil
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

func (l *LBD) MintNonFungible(contractId, tokenType, name, meta, to string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/mint", contractId, tokenType)

	r := MintNonFungibleRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
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

type MintMultipleNonFungibleRequest struct {
	*Request
	OwnerAddress string      `json:"ownerAddress"`
	OwnerSecret  string      `json:"ownerSecret"`
	MintList     []*MintList `json:"mintList"`
	ToUserId     string      `json:"toUserId,omitempty"`
	ToAddress    string      `json:"toAddress,omitempty"`
}

type MintList struct {
	TokenType string `json:"tokenType"`
	Name      string `json:"name"`
	Meta      string `json:"meta"`
}

func (r MintMultipleNonFungibleRequest) Encode() string {
	base := r.Request.Encode()
	names := make([]string, len(r.MintList))
	metas := make([]string, len(r.MintList))
	TokenTypes := make([]string, len(r.MintList))
	for i, m := range r.MintList {
		names[i] = m.Name
		metas[i] = m.Meta
		TokenTypes[i] = m.TokenType
	}
	mintList := fmt.Sprintf("mintList.meta=%s&mintList.name=%s&mintList.tokenType=%s",
		strings.Join(metas, ","),
		strings.Join(names, ","),
		strings.Join(TokenTypes, ","),
	)

	if r.ToUserId != "" {
		return fmt.Sprintf("%s?%s&ownerAddress=%s&ownerSecret=%s&toUserId=%s", base, mintList, r.OwnerAddress, r.OwnerSecret, r.ToUserId)
	}
	ret := fmt.Sprintf("%s?%s&ownerAddress=%s&ownerSecret=%s&toAddress=%s", base, mintList, r.OwnerAddress, r.OwnerSecret, r.ToAddress)
	fmt.Println(ret)
	return ret
}

func (l *LBD) MintMultipleNonFungible(contractId, to string, mintList []*MintList) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/multi-mint", contractId)

	r := MintMultipleNonFungibleRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
		MintList:     mintList,
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

func (l *LBD) UpdateNonFungibleInformation(contractId, tokenType, tokenIndex, name, meta string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s", contractId, tokenType, tokenIndex)

	r := UpdateNonFungibleInformationRequest{
		Request:      NewPutRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
		Name:         name,
		Meta:         meta,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}
