package lbd

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	CashewBaseURL = "https://test-api-blockchain.line.me"
)

var (
	rsLetterIdxBits       = 6
	rsLetterIdxMask int64 = 1<<rsLetterIdxBits - 1
	rsLetterIdxMax        = 63 / rsLetterIdxBits
	randSrc               = rand.NewSource(time.Now().UnixNano())
)

type LBD struct {
	baseURL   string
	apiKey    string
	apiSecret string
}

type Requester interface {
	Method() string
	Path() string
	Nonce() string
	Timestamp() string
	Encode() string
}

type Request struct {
	nonce     string
	timestamp int64
	method    string
	path      string
}

func NewGetRequest(path string) *Request {
	return NewRequest("GET", path)
}

func NewPostRequest(path string) *Request {
	return NewRequest("POST", path)
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

func (r *Request) Path() string {
	return r.path
}

func (r *Request) Nonce() string {
	return r.nonce
}

func (r *Request) Timestamp() string {
	return fmt.Sprint(r.timestamp)
}

func (r *Request) Encode() string {
	return fmt.Sprintf("%s%s%s%s", r.Nonce(), r.Timestamp(), r.method, r.path)
}

// type Method string

// const (
// 	MethodGet  Method = "GET"
// 	MethodPost        = "POST"
// 	MethodPut         = "PUT"
// )

type Response struct {
	ResponseTime  int64           `json:"responseTime"`
	StatusCode    int64           `json:"statusCode"`
	StatusMessage string          `json:"statusMessage"`
	ResponseData  json.RawMessage `json:"responseData"`
}

func NewLBD(apiKey string, secret string) (*LBD, error) {
	l := &LBD{
		baseURL:   CashewBaseURL,
		apiKey:    apiKey,
		apiSecret: secret,
	}
	return l, nil
}

func (l LBD) Sign(r Requester) string {
	msg := r.Encode()
	mac := hmac.New(sha512.New, []byte(l.apiSecret))
	mac.Write([]byte(msg))
	sig := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(sig)
}

func (l LBD) Do(r Requester, body []byte, sign bool) (*Response, error) {
	ctx := context.TODO()
	url := l.baseURL + r.Path()

	fmt.Println(url)

	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	if string(body) == "{}" {
		body = nil
	}

	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, r.Method(), url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("service-api-key", l.apiKey)
	req.Header.Add("Content-Type", "application/json")

	fmt.Println(string(body))
	fmt.Println(r.Encode())

	if sign {
		sig := l.Sign(r)
		req.Header.Add("nonce", r.Nonce())
		req.Header.Add("timestamp", r.Timestamp())
		req.Header.Add("signature", sig)
	}

	fmt.Println(req.Header)

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
		return ret, fmt.Errorf("Backend returns status %d msg: %s", ret.StatusCode, ret.StatusMessage)
	}

	return ret, nil
}

func NowMsec() int64 {
	return time.Now().UnixNano() / 1000000
}

func GenerateNonce(timestampMsec int64) string {
	// TODO The same nonce can’t be reused per service-api-key within 20 seconds.
	// An error is returned when the nonce of the successful request is reused within 20 seconds.
	return randString(8)
}

func randString(l int) string {
	rsLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, l)
	cache, remain := randSrc.Int63(), rsLetterIdxMax
	for i := l - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rsLetterIdxMax
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
