package lbd

import (
	"encoding/json"
	"fmt"
	"math/big"
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
