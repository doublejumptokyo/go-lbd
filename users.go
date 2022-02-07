package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type UserInformation struct {
	UserID        string `json:"userId"`
	WalletAddress string `json:"walletAddress"`
}

func (l LBD) RetrieveUserInformation(userId string) (*UserInformation, error) {
	r := NewGetRequest("/v1/users/" + userId)

	data, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(UserInformation)
	return ret, json.Unmarshal(data.ResponseData, ret)
}

func (l LBD) RetrieveUserWalletTransactionHistory(userId string) ([]*Transaction, error) {
	r := NewGetRequest(fmt.Sprintf("/v1/users/%s/transactions", userId))

	data, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := []*Transaction{}
	return ret, json.Unmarshal(data.ResponseData, &ret)
}

type NonFungible struct {
	Name          string `json:"name"`
	TokenType     string `json:"tokenType"`
	Meta          string `json:"meta"`
	NumberOfIndex string `json:"numberOfIndex"`
}

func (l LBD) RetrieveBalanceOfAllNonFungiblesUserWallet(userId, contractId string) ([]*NonFungible, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles", userId, contractId)

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

type NonFungibleToken struct {
	Name       string `json:"name"`
	TokenType  string `json:"tokenType"`
	TokenIndex string `json:"tokenIndex"`
	Meta       string `json:"meta"`
}

func (n *NonFungibleToken) ID() string {
	return fmt.Sprintf("%s%08s", n.TokenType, n.TokenIndex)
}

func (l LBD) RetrieveBalanceOfSpecificTypeOfNonFungiblesUserWallet(userId, contractId, tokenType string) ([]*NonFungibleToken, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/%s", userId, contractId, tokenType)

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
		all = append(all, ret...)
		page++
		if len(ret) < r.pager.Limit {
			break
		}
	}

	for _, t := range all {
		t.TokenType = tokenType
	}
	return all, nil
}

type SessionToken struct {
	RequestSessionToken string `json:"requestSessionToken"`
	RedirectURI         string `json:"redirectUri"`
}

func UnmarshalSessionToken(data []byte) (*SessionToken, error) {
	r := new(SessionToken)
	err := json.Unmarshal(data, r)
	return r, err
}

func (r *SessionToken) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IssueSessionTokenForBaseCoinTransferRequest struct {
	*Request
	ToUserId    string      `json:"toUserId,omitempty"`
	ToAddress   string      `json:"toAddress,omitempty"`
	Amount      string      `json:"amount"`
	RequestType RequestType `json:"-"`
	// LandingUri    string      `json:"landingUri,omitempty"`
}

func (r IssueSessionTokenForBaseCoinTransferRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s&amount=%s&toUserId=%s", base, r.Amount, r.ToUserId)
	}
	return fmt.Sprintf("%s&amount=%s&toAddress=%s", base, r.Amount, r.ToAddress)
}

func (l *LBD) IssueSessionTokenForBaseCoinTransfer(fromUserId, to string, amount *big.Int, requestType RequestType) (*SessionToken, error) {
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

type IssueSessionTokenForProxySettingRequest struct {
	*Request
	OwnerAddress string      `json:"ownerAddress"`
	OwnerSecret  string      `json:"ownerSecret"`
	RequestType  RequestType `json:"-"`
	// LandingUri    string      `json:"landingUri,omitempty"`
}

func (r IssueSessionTokenForProxySettingRequest) Encode() string {
	base := r.Request.Encode()
	return fmt.Sprintf("%s&ownerAddress=%s&ownerSecret=%s", base, r.OwnerAddress, r.OwnerSecret)
}

func (l *LBD) IssueSessionTokenForProxySetting(userId, contractId string, requestType RequestType) (*SessionToken, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/request-proxy?requestType=%s", userId, contractId, requestType)
	r := &IssueSessionTokenForProxySettingRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
	}

	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalSessionToken(resp.ResponseData)
}

type TransferNonFungibleUserWalletRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
}

func (r TransferNonFungibleUserWalletRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?ownerAddress=%s&ownerSecret=%s&toUserId=%s", base, r.OwnerAddress, r.OwnerSecret, r.ToUserId)
	}
	return fmt.Sprintf("%s?ownerAddress=%s&ownerSecret=%s&toAddress=%s", base, r.OwnerAddress, r.OwnerSecret, r.ToAddress)
}

func (l *LBD) TransferNonFungibleUserWallet(contractId, fromUserId, to, tokenType, tokenIndex string) (*Transaction, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/%s/%s/transfer", fromUserId, contractId, tokenType, tokenIndex)
	r := &TransferNonFungibleUserWalletRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
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

type SessionTokenStatus string

const (
	SessionTokenStatusUnknown      SessionTokenStatus = "Unknown"
	SessionTokenStatusAuthorized   SessionTokenStatus = "Authorized"
	SessionTokenStatusUnauthorized SessionTokenStatus = "Unauthorized"
)

func (l LBD) RetrieveSessionTokenStatus(requestSessionToken string) (SessionTokenStatus, error) {
	r := NewGetRequest("/v1/user-requests/" + requestSessionToken)

	data, err := l.Do(r, true)
	if err != nil {
		return SessionTokenStatusUnknown, err
	}

	ret := &struct {
		Status SessionTokenStatus `json:"status"`
	}{}

	return ret.Status, json.Unmarshal(data.ResponseData, ret)
}

func (l LBD) CommitTransaction(requestSessionToken string) (*Transaction, error) {
	r := NewPostRequest("/v1/user-requests/" + requestSessionToken + "/commit")
	data, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(data.ResponseData)
}

type Balance struct {
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
	Amount   string `json:"amount"`
}

func (l LBD) RetrieveBaseCoinBalanceUserWallet(userId string) (*Balance, error) {
	r := NewGetRequest(fmt.Sprintf("/v1/users/%s/base-coin", userId))
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(Balance)
	err = json.Unmarshal(resp.ResponseData, ret)
	return ret, err
}

type BalanceOfServiceTokens struct {
	ContractId string `json:"contractId"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	ImgUri     string `json:"imgUri"`
	Decimals   int64  `json:"decimals"`
	Amount     string `json:"amount"`
}

func (l LBD) RetrieveBalanceOfAllServiceTokensUserWallet(userId string) ([]*BalanceOfServiceTokens, error) {
	path := fmt.Sprintf("/v1/users/%s/service-tokens", userId)

	all := []*BalanceOfServiceTokens{}
	page := 1
	for {
		r := NewGetRequest(path)
		r.pager.Page = page
		r.pager.OrderBy = "asc"
		resp, err := l.Do(r, true)
		if err != nil {
			return nil, err
		}
		ret := []*BalanceOfServiceTokens{}
		err = json.Unmarshal(resp.ResponseData, &ret)
		if err != nil {
			return nil, err
		}
		all = append(all, ret...)
		page++
		if len(ret) < r.pager.Limit {
			break
		}
	}
	return all, nil
}

func (l LBD) RetrieveBalanceOfSpecificServiceTokenUserWallet(userId, contractId string) (*BalanceOfServiceTokens, error) {
	r := NewGetRequest(fmt.Sprintf("/v1/users/%s/service-tokens/%s", userId, contractId))
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}

	ret := new(BalanceOfServiceTokens)
	err = json.Unmarshal(resp.ResponseData, ret)
	return ret, err
}

type BalanceOfFungible struct {
	TokenType string `json:"tokenType"`
	Name      string `json:"name"`
	Meta      string `json:"meta"`
	Amount    string `json:"amount"`
}

func (l LBD) RetrieveBalanceOfAllFungiblesUserWallet(userId, contractId string) ([]*BalanceOfFungible, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/fungibles", userId, contractId)

	all := []*BalanceOfFungible{}
	page := 1
	for {
		r := NewGetRequest(path)
		r.pager.Page = page
		r.pager.OrderBy = "asc"
		resp, err := l.Do(r, true)
		if err != nil {
			return nil, err
		}
		ret := []*BalanceOfFungible{}
		err = json.Unmarshal(resp.ResponseData, &ret)
		if err != nil {
			return nil, err
		}
		all = append(all, ret...)
		page++
		if len(ret) < r.pager.Limit {
			break
		}
	}
	return all, nil
}

func (l LBD) RetrieveBalanceOfSpecificFungibleUserWallet(userId, contractId, tokenType string) (*BalanceOfFungible, error) {
	r := NewGetRequest(fmt.Sprintf("/v1/users/%s/item-tokens/%s/fungibles/%s", userId, contractId, tokenType))
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(BalanceOfFungible)
	err = json.Unmarshal(resp.ResponseData, ret)
	return ret, err
}

type BalanceOfNonFungiblesTokenType struct {
	List          []BalanceOfNonFungiblesTokenTypeList `json:"list"`
	PrePageToken  string                               `json:"prePageToken"`
	NextPageToken string                               `json:"nextPageToken"`
}

type BalanceOfNonFungiblesTokenTypeList struct {
	Type  BalanceOfNonFungiblesType  `json:"type"`
	Token BalanceOfNonFungiblesToken `json:"token"`
}

type BalanceOfNonFungiblesType struct {
	TokenType   string `json:"tokenType"`
	Name        string `json:"name"`
	Meta        string `json:"meta"`
	CreatedAt   int64  `json:"createdAt"`
	TotalSupply string `json:"totalSupply"`
	TotalMint   string `json:"totalMint"`
	TotalBurn   string `json:"totalBurn"`
}

type BalanceOfNonFungiblesToken struct {
	Name      string `json:"name"`
	TokenId   string `json:"tokenId"`
	CreatedAt int64  `json:"createdAt"`
	BurnedAt  int64  `json:"burnedAt"`
}

func (l LBD) RetrieveBalanceOfNonFungiblesWithTokenTypeUserWallet(userId, contractId, orderBy, pageToken string, limit uint32) (*BalanceOfNonFungiblesTokenType, error) {
	var r *Request
	if pageToken == "" {
		r = NewGetRequest(fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/with-type?limit=%d&orderBy=%s", userId, contractId, limit, orderBy))
	} else {
		r = NewGetRequest(fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/with-type?limit=%d&orderBy=%s&pageToken=%s", userId, contractId, limit, orderBy, pageToken))
	}
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(BalanceOfNonFungiblesTokenType)
	err = json.Unmarshal(resp.ResponseData, ret)
	return ret, err
}

type BalanceOfSpecificNonFungible struct {
	TokenIndex string `json:"tokenIndex"`
	Name       string `json:"name"`
	Meta       string `json:"meta"`
}

func (l LBD) RetrieveBalanceOfSpecificNonFungibleUserWallet(userId, contractId, tokenType, tokenIndex string) (*BalanceOfSpecificNonFungible, error) {
	r := NewGetRequest(fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/%s/%s", userId, contractId, tokenType, tokenIndex))
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(BalanceOfSpecificNonFungible)
	err = json.Unmarshal(resp.ResponseData, ret)
	fmt.Println(err)
	return ret, err
}

type ServiceTokenProxySet struct {
	IsApproved bool `json:"isApproved"`
}

func (l LBD) RetrieveWhetherTheServiceTokenProxySet(userId, contractId string) (*ServiceTokenProxySet, error) {
	r := NewGetRequest(fmt.Sprintf("/v1/users/%s/service-tokens/%s/proxy", userId, contractId))
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(ServiceTokenProxySet)
	err = json.Unmarshal(resp.ResponseData, ret)
	return ret, err
}

func (l LBD) RetrieveWhetherTheItemTokenProxySet(userId, contractId string) (*ServiceTokenProxySet, error) {
	r := NewGetRequest(fmt.Sprintf("/v1/users/%s/item-tokens/%s/proxy", userId, contractId))
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(ServiceTokenProxySet)
	err = json.Unmarshal(resp.ResponseData, ret)
	return ret, err
}

func (l LBD) IssueSessionTokenForServiceTokenTransfer(fromUserId, contractId, to string, amount *big.Int, requestType RequestType) (*SessionToken, error) {
	path := fmt.Sprintf("/v1/users/%s/service-tokens/%s/request-transfer?requestType=%s", fromUserId, contractId, requestType)
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

func (l LBD) IssueSessionTokenForServiceTokenProxySetting(userId, contractId string, requestType RequestType) (*SessionToken, error) {
	path := fmt.Sprintf("/v1/users/%s/service-tokens/%s/request-proxy?requestType=%s", userId, contractId, requestType)
	r := &IssueSessionTokenForProxySettingRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
	}
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalSessionToken(resp.ResponseData)
}

type Transfer struct {
	TxHasn string `json:"txHash"`
}

func UnmarshalTransfer(data []byte) (*Transfer, error) {
	r := new(Transfer)
	err := json.Unmarshal(data, r)
	return r, err
}

func (r *Transfer) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TransferDelegatedRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	Amount       string `json:"amount"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
}

func (r TransferDelegatedRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?amount=%s&ownerAddress=%s&ownerSecret=%s&toUserId=%s", base, r.Amount, r.OwnerAddress, r.OwnerSecret, r.ToUserId)
	}
	return fmt.Sprintf("%s?amount=%s&ownerAddress=%s&ownerSecret=%s&toAddress=%s", base, r.Amount, r.OwnerAddress, r.OwnerSecret, r.ToAddress)
}

func (l LBD) TransferDelegatedServiceTokenUserWallet(userId, contractId, to string, amount *big.Int) (*Transfer, error) {
	path := fmt.Sprintf("/v1/users/%s/service-tokens/%s/transfer", userId, contractId)
	r := TransferDelegatedRequest{
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

	return UnmarshalTransfer(resp.ResponseData)
}

func (l LBD) TransferDelegatedFungibleUserWallet(userId, contractId, tokenType, to string, amount *big.Int) (*Transfer, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/fungibles/%s/transfer", userId, contractId, tokenType)
	r := TransferDelegatedRequest{
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
	return UnmarshalTransfer(resp.ResponseData)
}

type TransferDelegatedNonFungibleRequest struct {
	*Request
	OwnerAddress string `json:"ownerAddress"`
	OwnerSecret  string `json:"ownerSecret"`
	ToUserId     string `json:"toUserId,omitempty"`
	ToAddress    string `json:"toAddress,omitempty"`
}

func (r TransferDelegatedNonFungibleRequest) Encode() string {
	base := r.Request.Encode()
	if r.ToUserId != "" {
		return fmt.Sprintf("%s?ownerAddress=%s&ownerSecret=%s&toUserId=%s", base, r.OwnerAddress, r.OwnerSecret, r.ToUserId)
	}
	return fmt.Sprintf("%s?ownerAddress=%s&ownerSecret=%s&toAddress=%s", base, r.OwnerAddress, r.OwnerSecret, r.ToAddress)
}

func (l LBD) TransferDelegatedNonFungible(userId, contractId, tokenType, tokenIndex, to string) (*Transfer, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/%s/%s/transfer", userId, contractId, tokenType, tokenIndex)
	r := &TransferDelegatedNonFungibleRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
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
	return UnmarshalTransfer(resp.ResponseData)
}

type TransferDelegatedNonFungiblesRequest struct {
	*Request
	OwnerAddress string         `json:"ownerAddress"`
	OwnerSecret  string         `json:"ownerSecret"`
	ToUserId     string         `json:"toUserId,omitempty"`
	ToAddress    string         `json:"toAddress,omitempty"`
	TransferList []TransferList `json:"transferList"`
}

func (r TransferDelegatedNonFungiblesRequest) Encode() string {
	base := r.Request.Encode()
	tokenIds := make([]string, len(r.TransferList))

	for i, m := range r.TransferList {
		tokenIds[i] = m.TokenId
	}
	transferList := fmt.Sprintf("transferList.tokenId=%s", strings.Join(tokenIds, ","))

	if r.ToUserId != "" {
		return fmt.Sprintf("%s?ownerAddress=%s&ownerSecret=%s&toUserId=%s&%s", base, r.OwnerAddress, r.OwnerSecret, r.ToUserId, transferList)
	}
	return fmt.Sprintf("%s?ownerAddress=%s&ownerSecret=%s&toAddress=%s&%s", base, r.OwnerAddress, r.OwnerSecret, r.ToAddress, transferList)
}

func (l LBD) BatchTransferDelegatedNonFungiblesUserWallet(userId, contractId, to string, transferList []TransferList) (*Transfer, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/batch-transfer", userId, contractId)
	r := &TransferDelegatedNonFungiblesRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: l.Owner.Address,
		OwnerSecret:  l.Owner.Secret,
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
	return UnmarshalTransfer(resp.ResponseData)
}
