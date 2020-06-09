package lbd

import (
	"encoding/json"
	"fmt"
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

func UnmarshalSessionToken(data []byte) (*SessionToken, error) {
	r := new(SessionToken)
	err := json.Unmarshal(data, r)
	return r, err
}

func (r *SessionToken) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SessionToken struct {
	RequestSessionToken string `json:"requestSessionToken"`
	RedirectURI         string `json:"redirectUri"`
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

func (l *LBD) IssueSessionTokenForProxySetting(userId, contractId string, requestType RequestType, owner *Wallet) (*SessionToken, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/request-proxy?requestType=%s", userId, contractId, requestType)
	r := &IssueSessionTokenForProxySettingRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: owner.Address,
		OwnerSecret:  owner.Secret,
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

func (l *LBD) TransferNonFungibleUserWallet(contractId, fromUserId, to, tokenType, tokenIndex string, owner *Wallet) (*Transaction, error) {
	path := fmt.Sprintf("/v1/users/%s/item-tokens/%s/non-fungibles/%s/%s/transfer", fromUserId, contractId, tokenType, tokenIndex)
	r := &TransferNonFungibleUserWalletRequest{
		Request:      NewPostRequest(path),
		OwnerAddress: owner.Address,
		OwnerSecret:  owner.Secret,
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
	SessionTokenStatusAuthorized                      = "Authorized"
	SessionTokenStatusUnauthorized                    = "Unauthorized"
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
