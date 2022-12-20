package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type Meta struct {
	data map[string]string
}

func NewMeta() *Meta {
	return &Meta{
		data: map[string]string{},
	}
}

func (m *Meta) Set(key, value string) (err error) {
	if len(key) < 1 && 15 < len(key) {
		return fmt.Errorf("invalid key length")
	}
	if len(value) < 1 && 15 < len(value) {
		return fmt.Errorf("invalid value length")
	}
	m.data[key] = value
	return nil
}

func (m *Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

func UnmarshalMeta(data []byte) (*Meta, error) {
	m := new(Meta)
	return m, json.Unmarshal(data, &m.data)
}

func (m *Meta) String() string {
	b, _ := m.MarshalJSON()
	return string(b)
}

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

func (l LBD) ListAllFungibles(contractId string, pager *Pager) ([]*TokenType, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles", contractId)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*TokenType{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

type FungibleInformation struct {
	TokenType   string `json:"tokenType"`
	Name        string `json:"name"`
	Meta        string `json:"meta"`
	CreatedAt   int64  `json:"createdAt"`
	TotalSupply string `json:"totalSupply"`
	TotalMint   string `json:"totalMint"`
	TotalBurn   string `json:"totalBurn"`
}

func (l LBD) RetrieveFungibleInformation(contractId, tokenType string) (*FungibleInformation, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/%s", contractId, tokenType)
	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(FungibleInformation)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

type FungibleHolers struct {
	WalletAddress string `json:"walletAddress"`
	UserID        string `json:"userId"`
	Amount        string `json:"amount"`
}

func (l LBD) RetrieveAllFungibleHolders(contractId, tokenType string, pager *Pager) ([]*FungibleHolers, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/%s/holders", contractId, tokenType)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*FungibleHolers{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

func (l LBD) ListAllNonFungibles(contractId string, pager *Pager) ([]*TokenType, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles", contractId)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*TokenType{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, err
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

func (l *LBD) RetrieveNonFungibleTokenType(contractId, tokenType string, pager *Pager) (*TokenType, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s", contractId, tokenType)
	if pager == nil {
		pager = &Pager{
			Limit:   10,
			Page:    1,
			OrderBy: "desc",
		}
	}

	r := NewGetRequest(path)
	r.pager = pager
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(TokenType)
	return ret, json.Unmarshal(resp.ResponseData, ret)
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

// Holders Response Struct
type Holder struct {
	WalletAddress *string `json:"walletAddress"`
	UserID        *string `json:"userId"`
	NumberOfIndex string  `json:"numberOfIndex"`
}

func (l LBD) RetrieveHolderOfSpecificNonFungible(contractId, tokenType string, pager *Pager) ([]*Holder, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/holders", contractId, tokenType)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*Holder{}
	err = json.Unmarshal(resp.ResponseData, &ret)

	return ret, err
}

type ItemTokenHolder struct {
	WalletAddress *string `json:"walletAddress"`
	UserID        *string `json:"userId"`
	TokenID       *string `json:"tokenId"`
	Amount        string  `json:"amount"`
}

func (l LBD) RetrieveTheHolderOfSpecificNonFungible(contractId, tokenType, tokenIndex string) (*ItemTokenHolder, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s/holder", contractId, tokenType, tokenIndex)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(ItemTokenHolder)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

func (l LBD) ListTheChildrenOfNonFungible(contractId, tokenType, tokenIndex string, pager *Pager) ([]*NonFungibleInformation, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s/children", contractId, tokenType, tokenIndex)

	r := NewGetRequest(path)
	r.pager = pager

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*NonFungibleInformation{}
	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

type ParentNonFungible struct {
	Name      string `json:"name"`
	TokenId   string `json:"tokenId"`
	Meta      string `json:"meta"`
	CreatedAt int64  `json:"createdAt"`
	BurnedAt  int64  `json:"burnedAt"`
}

func (l LBD) RetrieveTheParentOfNonFungible(contractId, tokenType, tokenIndex string) (*ParentNonFungible, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s/parent", contractId, tokenType, tokenIndex)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(ParentNonFungible)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

func (l LBD) RetrieveTheRootOfNonFungible(contractId, tokenType, tokenIndex string) (*ParentNonFungible, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s/root", contractId, tokenType, tokenIndex)

	r := NewGetRequest(path)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(ParentNonFungible)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}

type FungibleTokenResponse struct {
	TokenType    string `json:"tokenType"`
	Url          string `json:"url"`
	Status       string `json:"status"`
	DetailStatus string `json:"detailStatus"`
}

// Deprecated: Use RetrieveFungibleTokenMediaResourceStatus or RetrieveFungibleTokenThumbnailsStatus instead.
func (l LBD) RetrieveTheStatusOfMultipleFungibleTokenIcons(contractId, requestId string) ([]*FungibleTokenResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/icon/%s/status", contractId, requestId)

	r := NewGetRequest(path)

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*FungibleTokenResponse{}

	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type NonFungibleTokenResponse struct {
	TokenType    string `json:"tokenType"`
	TokenIndex   string `json:"tokenIndex"`
	Url          string `json:"url"`
	Status       string `json:"status"`
	DetailStatus string `json:"detailStatus"`
}

// Deprecated: Use RetrieveNonFungibleTokenMediaResourceStatus or RetrieveNonFungibleTokenThumbnailsStatus instead.
func (l LBD) RetrieveTheStatusOfMultipleNonFungibleTokenIcons(contractId, requestId string) ([]*NonFungibleTokenResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/icon/%s/status", contractId, requestId)

	r := NewGetRequest(path)

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*NonFungibleTokenResponse{}

	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
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

type UpdateMultipleFungibleTokenIconsRequest struct {
	*Request
	UpdateList []*UpdateFungibleList `json:"updateList"`
}
type UpdateMultipleFungibleTokenUpdateListRequest struct {
	*Request
	UpdateList []*UpdateFungibleList `json:"updateList"`
}

type UpdateList struct {
	TokenType  string `json:"tokenType"`
	TokenIndex string `json:"tokenIndex"`
}

type UpdateFungibleList struct {
	TokenType string `json:"tokenType"`
}

func (r UpdateMultipleFungibleTokenIconsRequest) Encode() string {
	base := r.Request.Encode()
	types := make([]string, len(r.UpdateList))

	for i, m := range r.UpdateList {
		types[i] = m.TokenType
	}
	updateList := fmt.Sprintf("updateList.tokenType=%s",
		strings.Join(types, ","),
	)

	ret := fmt.Sprintf("%s?%s", base, updateList)
	return ret
}

func (r UpdateMultipleFungibleTokenUpdateListRequest) Encode() string {
	base := r.Request.Encode()
	types := make([]string, len(r.UpdateList))

	for i, m := range r.UpdateList {
		types[i] = m.TokenType
	}
	updateList := fmt.Sprintf("updateList.tokenType=%s",
		strings.Join(types, ","),
	)

	ret := fmt.Sprintf("%s?%s", base, updateList)
	return ret
}

func (r UpdateMultipleNonFungibleTokenIconsRequest) Encode() string {
	base := r.Request.Encode()
	types := make([]string, len(r.UpdateList))
	indexes := make([]string, len(r.UpdateList))

	for i, m := range r.UpdateList {
		types[i] = m.TokenType
		indexes[i] = m.TokenIndex
	}
	updateList := fmt.Sprintf("updateList.tokenIndex=%s&updateList.tokenType=%s",
		strings.Join(indexes, ","),
		strings.Join(types, ","),
	)

	ret := fmt.Sprintf("%s?%s", base, updateList)
	return ret
}

func (r UpdateMultipleNonFungibleTokenUpdateListRequest) Encode() string {
	base := r.Request.Encode()
	types := make([]string, len(r.UpdateList))
	indexes := make([]string, len(r.UpdateList))

	for i, m := range r.UpdateList {
		types[i] = m.TokenType
		indexes[i] = m.TokenIndex
	}
	updateList := fmt.Sprintf("updateList.tokenIndex=%s&updateList.tokenType=%s",
		strings.Join(indexes, ","),
		strings.Join(types, ","),
	)

	ret := fmt.Sprintf("%s?%s", base, updateList)
	return ret
}

type UpdateMultipleNonFungibleTokenIconsRequest struct {
	*Request
	UpdateList []*UpdateList `json:"updateList"`
}

type UpdateMultipleNonFungibleTokenUpdateListRequest struct {
	*Request
	UpdateList []*UpdateList `json:"updateList"`
}

type UpdateMultipleTokenIconsResponse struct {
	RequestId string `json:"requestId"`
}

type UpdateMediaResourcesResponse struct {
	RequestId string `json:"requestId"`
}

type UpdateThumbnailsResponse struct {
	RequestId string `json:"requestId"`
}

// Deprecated: Use UpdateNonFungibleTokenThumbnails or UpdateNonFungibleTokenMediaResources or  instead.
func (l *LBD) UpdateMultipleNonFungibleTokenIcons(contactId string, updateList []*UpdateList) (*UpdateMultipleTokenIconsResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/icon", contactId)

	r := UpdateMultipleNonFungibleTokenIconsRequest{
		Request:    NewPutRequest(path),
		UpdateList: updateList,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UpdateMultipleTokenIconsResponse)
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

// Deprecated: Use UpdateFungibleTokenThumbnails or UpdateFungibleTokenMediaResources instead.
func (l *LBD) UpdateMultipleFungibleTokenIcons(contactId string, updateList []*UpdateFungibleList) (*UpdateMultipleTokenIconsResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/icon", contactId)

	r := UpdateMultipleFungibleTokenIconsRequest{
		Request:    NewPutRequest(path),
		UpdateList: updateList,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UpdateMultipleTokenIconsResponse)
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

func (r UpdateFungibleInformationRequest) Encode() string {
	base := r.Request.Encode()
	if r.Name != "" && r.Meta != "" {
		return fmt.Sprintf("%s?meta=%s&name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Meta, r.Name, r.OwnerAddress, r.OwnerSecret)
	}
	if r.Name != "" {
		return fmt.Sprintf("%s?name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Name, r.OwnerAddress, r.OwnerSecret)
	}
	return fmt.Sprintf("%s?meta=%s&ownerAddress=%s&ownerSecret=%s", base, r.Meta, r.OwnerAddress, r.OwnerSecret)
}

type UpdateFungibleInformationRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Name         string `json:"name,omitempty"`
	Meta         string `json:"meta,omitempty"`
}

func (l *LBD) UpdateFungibleInformation(contractId, tokenType, name, meta string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/%s", contractId, tokenType)

	r := UpdateFungibleInformationRequest{
		Request:      NewPutRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
	}

	if name != "" {
		r.Name = name
	}

	if meta != "" {
		r.Meta = meta
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

func (r AttachNonFungibleRequest) Encode() string {
	base := r.Request.Encode()
	if r.TokenHolderUserId != "" {
		return fmt.Sprintf("%s?parentTokenId=%s&serviceWalletAddress=%s&serviceWalletSecret=%s&tokenHolderUserId=%s", base, r.ParentTokenId, r.ServiceWalletAddress, r.ServiceWalletSecret, r.TokenHolderUserId)
	}
	return fmt.Sprintf("%s?parentTokenId=%s&serviceWalletAddress=%s&serviceWalletSecret=%s&tokenHolderAddress=%s", base, r.ParentTokenId, r.ServiceWalletAddress, r.ServiceWalletSecret, r.TokenHolderAddress)
}

type AttachNonFungibleRequest struct {
	*Request
	ParentTokenId        string `json:"parentTokenId"`
	ServiceWalletAddress string `json:"serviceWalletAddress"`
	ServiceWalletSecret  string `json:"serviceWalletSecret"`
	TokenHolderAddress   string `json:"tokenHolderAddress,omitempty"`
	TokenHolderUserId    string `json:"tokenHolderUserId,omitempty"`
}

func (l *LBD) AttachNonFungibleAnother(contractId, tokenType, tokenIndex, parentTokenId, to string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s/parent", contractId, tokenType, tokenIndex)

	r := AttachNonFungibleRequest{
		Request:              NewPostRequest(path),
		ParentTokenId:        parentTokenId,
		ServiceWalletAddress: l.Owner.Address,
		ServiceWalletSecret:  l.Owner.Secret,
	}

	if l.IsAddress(to) {
		r.TokenHolderAddress = to
	} else {
		r.TokenHolderUserId = to
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	return UnmarshalTransaction(resp.ResponseData)
}

func (r DetachNonFungibleParentRequest) Encode() string {
	base := r.Request.Encode()
	if r.TokenHolderUserId != "" {
		return fmt.Sprintf("%s?serviceWalletAddress=%s&serviceWalletSecret=%s&tokenHolderUserId=%s", base, r.ServiceWalletAddress, r.ServiceWalletSecret, r.TokenHolderUserId)
	}
	return fmt.Sprintf("%s?serviceWalletAddress=%s&serviceWalletSecret=%s&tokenHolderAddress=%s", base, r.ServiceWalletAddress, r.ServiceWalletSecret, r.TokenHolderAddress)
}

type DetachNonFungibleParentRequest struct {
	*Request
	ServiceWalletAddress string `json:"serviceWalletAddress"`
	ServiceWalletSecret  string `json:"serviceWalletSecret"`
	TokenHolderAddress   string `json:"tokenHolderAddress,omitempty"`
	TokenHolderUserId    string `json:"tokenHolderUserId,omitempty"`
}

func (l *LBD) DetachNonFungibleParent(contractId, tokenType, tokenIndex, to string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s/parent", contractId, tokenType, tokenIndex)

	r := DetachNonFungibleParentRequest{
		Request:              NewDeleteRequest(path),
		ServiceWalletAddress: l.Owner.Address,
		ServiceWalletSecret:  l.Owner.Secret,
	}

	if l.IsAddress(to) {
		r.TokenHolderAddress = to
	} else {
		r.TokenHolderUserId = to
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	return UnmarshalTransaction(resp.ResponseData)
}

// Mint or Burn

func (r CreateFungibleRequest) Encode() string {
	base := r.Request.Encode()
	if r.Meta != "" {
		return fmt.Sprintf("%s?meta=%s&name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Meta, r.Name, r.OwnerAddress, r.OwnerSecret)
	}
	return fmt.Sprintf("%s?name=%s&ownerAddress=%s&ownerSecret=%s", base, r.Name, r.OwnerAddress, r.OwnerSecret)
}

type CreateFungibleRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Name         string `json:"name"`
	Meta         string `json:"meta,omitempty"`
}

func (l *LBD) IssueFungible(contractId, name, meta string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles", contractId)

	r := CreateFungibleRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
		Name:         name,
	}

	if meta != "" {
		r.Meta = meta
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

func (r MintFungibleRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?amount=%s&ownerAddress=%s&ownerSecret=%s&toUserId=%s", base, r.Amount, r.OwnerAddress, r.OwnerSecret, r.ToUserId)
	}
	return fmt.Sprintf("%s?amount=%s&ownerAddress=%s&ownerSecret=%s&toAddress=%s", base, r.Amount, r.OwnerAddress, r.OwnerSecret, r.ToAddress)
}

type MintFungibleRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
	Amount       string `json:"amount"`
}

func (l *LBD) MintFungible(contractId, tokenType, to string, amount *big.Int) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/%s/mint", contractId, tokenType)
	r := MintFungibleRequest{
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

func (r BurnItemTokenRequest) Encode() string {
	base := r.Request.Encode()
	if r.FromUserId != "" {
		return fmt.Sprintf("%s?amount=%s&fromUserId=%s&ownerAddress=%s&ownerSecret=%s", base, r.Amount, r.FromUserId, r.OwnerAddress, r.OwnerSecret)
	}
	return fmt.Sprintf("%s?amount=%s&fromAddress=%s&ownerAddress=%s&ownerSecret=%s", base, r.Amount, r.FromAddress, r.OwnerAddress, r.OwnerSecret)
}

type BurnItemTokenRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Amount       string `json:"amount,omitempty"`
	FromUserId   string `json:"fromUserId,omitempty"`
	FromAddress  string `json:"fromAddress,omitempty"`
}

func (l *LBD) BurnFungible(contractId, tokenType, from string, amount *big.Int) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/%s/burn", contractId, tokenType)

	r := BurnItemTokenRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
		Amount:       amount.String(),
	}

	if l.IsAddress(from) {
		r.FromAddress = from
	} else {
		r.FromUserId = from
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

func encodeValidates(checkList []string, key string) string {
	var isEmpty bool
	var result string

	for _, r := range checkList {
		if r != "" {
			isEmpty = true
			break
		}
	}

	if isEmpty {
		result = fmt.Sprintf(key, strings.Join(checkList, ","))
	}
	return result
}

func (r MintMultipleNonFungibleRecipientsRequest) Encode() string {
	base := r.Request.Encode()
	tokenTypes := make([]string, len(r.MintList))
	names := make([]string, len(r.MintList))
	metas := make([]string, len(r.MintList))
	toUserIds := make([]string, len(r.MintList))
	toAddresses := make([]string, len(r.MintList))

	for i, m := range r.MintList {
		metas[i] = m.Meta
		names[i] = m.Name
		toAddresses[i] = m.ToAddress
		tokenTypes[i] = m.TokenType
		toUserIds[i] = m.ToUserId
	}

	meta := encodeValidates(metas, "mintList.meta=%s&")
	toAddress := encodeValidates(toAddresses, "mintList.toAddress=%s&")
	toUserId := encodeValidates(toUserIds, "mintList.toUserId=%s&")

	name := fmt.Sprintf("mintList.name=%s&", strings.Join(names, ","))
	tokenType := fmt.Sprintf("mintList.tokenType=%s&", strings.Join(tokenTypes, ","))

	mintList := fmt.Sprintf("%s%s%s%s%s", meta, name, toAddress, toUserId, tokenType)

	return fmt.Sprintf("%s?%sownerAddress=%s&ownerSecret=%s", base, mintList, r.OwnerAddress, r.OwnerSecret)
}

type MultiMintList struct {
	TokenType string `json:"tokenType"`
	Name      string `json:"name"`
	Meta      string `json:"meta,omitempty"`
	ToAddress string `json:"toAddress,omitempty"`
	ToUserId  string `json:"toUserId,omitempty"`
}

type MintMultipleNonFungibleRecipientsRequest struct {
	*Request
	OwnerAddress string           `json:"ownerAddress"`
	OwnerSecret  string           `json:"ownerSecret"`
	MintList     []*MultiMintList `json:"mintList"`
}

func (l *LBD) MintMultipleNonFungibleRecipients(contractId string, mintList []*MultiMintList) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/multi-recipients/multi-mint", contractId)

	r := MintMultipleNonFungibleRecipientsRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
		MintList:     mintList,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

type NonFungibleTokenBurnRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	FromUserId   string `json:"fromUserId,omitempty"`
	FromAddress  string `json:"fromAddress,omitempty"`
}

func (r NonFungibleTokenBurnRequest) Encode() string {
	base := r.Request.Encode()
	if r.FromUserId != "" {
		return fmt.Sprintf("%s?fromUserId=%s&ownerAddress=%s&ownerSecret=%s", base, r.FromUserId, r.OwnerAddress, r.OwnerSecret)
	}
	return fmt.Sprintf("%s?fromAddress=%s&ownerAddress=%s&ownerSecret=%s", base, r.FromAddress, r.OwnerAddress, r.OwnerSecret)
}

func (l *LBD) BurnNonFungible(contractId, tokenType, tokenIndex, from string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/%s/%s/burn", contractId, tokenType, tokenIndex)

	r := NonFungibleTokenBurnRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
	}

	if l.IsAddress(from) {
		r.FromAddress = from
	} else {
		r.FromUserId = from
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

// Get media resource status for multiple fungible tokens
func (l LBD) RetrieveFungibleTokenMediaResourceStatus(contractId, requestId string) ([]*FungibleTokenResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/media-resources/%s/status", contractId, requestId)

	r := NewGetRequest(path)

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*FungibleTokenResponse{}

	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Get thumbnail status of multiple fungible tokens
func (l LBD) RetrieveFungibleTokenThumbnailStatus(contractId, requestId string) ([]*FungibleTokenResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/thumbnails/%s/status", contractId, requestId)

	r := NewGetRequest(path)

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*FungibleTokenResponse{}

	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Get media resource status for multiple non-fungible tokens
func (l LBD) RetrieveNonFungibleTokenMediaResourceStatus(contractId, requestId string) ([]*NonFungibleTokenResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/media-resources/%s/status", contractId, requestId)

	r := NewGetRequest(path)

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*NonFungibleTokenResponse{}

	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Get thumbnail status of multiple non-fungible tokens
func (l LBD) RetrieveNonFungibleTokenThumbnailStatus(contractId, requestId string) ([]*NonFungibleTokenResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/thumbnails/%s/status", contractId, requestId)

	r := NewGetRequest(path)

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*NonFungibleTokenResponse{}

	err = json.Unmarshal(resp.ResponseData, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Update media resources for multiple fungible tokens
func (l *LBD) UpdateFungibleTokenMediaResources(contactId string, updateList []*UpdateFungibleList) (*UpdateMediaResourcesResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/media-resources", contactId)

	r := UpdateMultipleFungibleTokenUpdateListRequest{
		Request:    NewPutRequest(path),
		UpdateList: updateList,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UpdateMediaResourcesResponse)
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

// Update thumbnails for multiple fungible tokens
func (l *LBD) UpdateFungibleTokenThumbnails(contactId string, updateList []*UpdateFungibleList) (*UpdateThumbnailsResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/fungibles/thumbnails", contactId)

	r := UpdateMultipleFungibleTokenUpdateListRequest{
		Request:    NewPutRequest(path),
		UpdateList: updateList,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UpdateThumbnailsResponse)
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

// Update media resources for multiple non-fungible tokens
func (l *LBD) UpdateNonFungibleTokenMediaResources(contactId string, updateList []*UpdateList) (*UpdateMediaResourcesResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/media-resources", contactId)

	r := UpdateMultipleNonFungibleTokenUpdateListRequest{
		Request:    NewPutRequest(path),
		UpdateList: updateList,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UpdateMediaResourcesResponse)
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}

// Update thumbnails for multiple non-fungible tokens
func (l *LBD) UpdateNonFungibleTokenThumbnails(contactId string, updateList []*UpdateList) (*UpdateThumbnailsResponse, error) {
	path := fmt.Sprintf("/v1/item-tokens/%s/non-fungibles/thumbnails", contactId)

	r := UpdateMultipleNonFungibleTokenUpdateListRequest{
		Request:    NewPutRequest(path),
		UpdateList: updateList,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UpdateThumbnailsResponse)
	return ret, json.Unmarshal(resp.ResponseData, &ret)
}
