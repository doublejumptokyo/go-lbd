package lbd

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	CashewBaseURL = "https://test-api.blockchain.line.me"
	DaphneBaseURL = "https://api.blockchain.line.me"
)

type Network string

const (
	Cashew Network = "Cashew"
	Daphne Network = "Daphne"
)

var (
	rsLetterIdxBits       = 6
	rsLetterIdxMask int64 = 1<<rsLetterIdxBits - 1
	rsLetterIdxMax        = 63 / rsLetterIdxBits
)

type RequestType string

const (
	RequestTypeRedirectUri RequestType = "redirectUri"
	RequestTypeAOA         RequestType = "aoa"
)

const (
	DefaultLimit int = 50
)

type LBD struct {
	Network   Network
	baseURL   string
	apiKey    string
	apiSecret string
	Owner     *Wallet
	Debug     bool
}

func NewLBD(network Network, url string, apiKey string, secret string, owner *Wallet) (*LBD, error) {
	l := &LBD{
		Network:   network,
		baseURL:   url,
		apiKey:    apiKey,
		apiSecret: secret,
		Owner:     owner,
		Debug:     false,
	}
	return l, nil
}

func NewCashew(apiKey string, secret string, owner *Wallet) (*LBD, error) {
	l, err := NewLBD(
		Cashew,
		CashewBaseURL,
		apiKey,
		secret,
		owner,
	)
	return l, err
}

func NewDaphne(apiKey string, secret string, owner *Wallet) (*LBD, error) {
	l, err := NewLBD(
		Daphne,
		DaphneBaseURL,
		apiKey,
		secret,
		owner,
	)
	return l, err
}

func (l LBD) Sign(r Requester) string {
	msg := r.Encode()
	mac := hmac.New(sha512.New, []byte(l.apiSecret))
	_, _ = mac.Write([]byte(msg))
	sig := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(sig)
}

func (l LBD) IsAddress(s string) bool {
	prefix := "link"
	if l.Network == Cashew {
		prefix = "tlink"
	}
	return strings.HasPrefix(s, prefix)
}

func (l *LBD) Do(r Requester, sign bool) (*Response, error) {
	ctx := context.TODO()
	url := l.baseURL + r.URI()

	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	if string(body) == "{}" {
		body = nil
	}

	client := newClient()
	req, err := http.NewRequestWithContext(ctx, r.Method(), url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("service-api-key", l.apiKey)
	req.Header.Add("Content-Type", "application/json")

	if r.HasQuery() {
		req.URL.RawQuery = r.Query().Encode()
	}

	if sign {
		sig := l.Sign(r)
		req.Header.Add("nonce", r.Nonce())
		req.Header.Add("timestamp", r.Timestamp())
		req.Header.Add("signature", sig)
	}

	if l.Debug {
		fmt.Println(url)
		fmt.Println(string(body))
		fmt.Println(r.Encode())
		fmt.Println(req.Header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := new(Response)
	err = json.Unmarshal(buf, ret)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return ret, fmt.Errorf("LBD: Backend returns status: %d msg: %s", ret.StatusCode, ret.StatusMessage)
	}

	return ret, nil
}

func newClient() *http.Client {
	return new(http.Client)
}

type Requester interface {
	Method() string
	URI() string
	Path() string
	Nonce() string
	Timestamp() string
	Encode() string
	HasQuery() bool
	Query() url.Values
	RawQuery() string
}

type Request struct {
	nonce     string
	timestamp int64
	method    string
	path      string
	pager     *Pager
	query     *url.Values
}

type Pager struct {
	Limit   int
	Page    int
	OrderBy string
}

func NewGetRequest(path string) *Request {
	req := NewRequest("GET", path)
	req.pager = &Pager{
		Limit:   DefaultLimit,
		Page:    1,
		OrderBy: "desc",
	}
	return req
}

func NewGetRequestWithQuery(path string, query url.Values) *Request {
	req := NewRequest("GET", path)
	req.query = &query
	return req
}

func NewPostRequest(path string) *Request {
	return NewRequest("POST", path)
}

func NewPutRequest(path string) *Request {
	return NewRequest("PUT", path)
}

func NewDeleteRequest(path string) *Request {
	return NewRequest("DELETE", path)
}

func NewRequest(method, path string) *Request {
	now := NowMsec()
	return &Request{
		nonce:     GenerateNonce(now),
		timestamp: now,
		method:    method,
		path:      path,
	}
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) URI() string {
	if r.pager != nil {
		return fmt.Sprintf("%s?limit=%d&orderBy=%s&page=%d", r.Path(), r.pager.Limit, r.pager.OrderBy, r.pager.Page)
	}

	return r.Path()
}

func (r *Request) Path() string {
	return r.path
}

func (r *Request) HasQuery() bool {
	return r.query != nil
}

func (r *Request) Query() url.Values {
	if r.HasQuery() {
		return *r.query
	}
	return nil
}

func (r *Request) RawQuery() string {
	q := r.Query()
	keys := make([]string, 0, len(q))
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var params []string
	for _, k := range keys {
		qs := q[k]
		for _, q := range qs {
			params = append(params, fmt.Sprintf("%s=%s", k, q))
		}
	}
	return strings.Join(params, "&")
}

func (r *Request) Nonce() string {
	return r.nonce
}

func (r *Request) Timestamp() string {
	return fmt.Sprint(r.timestamp)
}

func (r *Request) Encode() string {
	if r.HasQuery() {
		// query string must not be URL encoded when signing
		return fmt.Sprintf("%s%s%s%s?%s", r.Nonce(), r.Timestamp(), r.method, r.URI(), r.RawQuery())
	}
	return fmt.Sprintf("%s%s%s%s", r.Nonce(), r.Timestamp(), r.method, r.URI())
}

type Response struct {
	ResponseTime  int64           `json:"responseTime"`
	StatusCode    int64           `json:"statusCode"`
	StatusMessage string          `json:"statusMessage"`
	ResponseData  json.RawMessage `json:"responseData"`
}

func NowMsec() int64 {
	return time.Now().UnixNano() / 1000000
}

func GenerateNonce(timestampMsec int64) string {
	// TODO The same nonce can’t be reused per service-api-key within 20 seconds.
	// An error is returned when the nonce of the successful request is reused within 20 seconds.
	return randString(8)
}

func randInt63() int64 {
	b := [8]byte{}
	ct, _ := rand.Read(b[:])
	return (int64)(binary.BigEndian.Uint64(b[:ct]) >> 1)
}

func randString(l int) string {
	rsLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, l)
	cache, remain := randInt63(), rsLetterIdxMax
	for i := l - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randInt63(), rsLetterIdxMax
		}
		idx := int(cache & rsLetterIdxMask)
		if idx < len(rsLetters) {
			b[i] = rsLetters[idx]
			i--
		}
		cache >>= rsLetterIdxBits
		remain--
	}
	return string(b)
}
