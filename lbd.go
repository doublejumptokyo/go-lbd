package lbd

import (
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

type Response struct {
	ResponseTime  int64       `json:"responseTime"`
	StatusCode    int64       `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	ResponseData  interface{} `json:"responseData"`
}

func NewLBD(apiKey string, secret string) (*LBD, error) {
	l := &LBD{
		baseURL:   CashewBaseURL,
		apiKey:    apiKey,
		apiSecret: secret,
	}
	return l, nil
}

func (l LBD) Sign(nonce string, timestampMsec int64, method, path, query string) string {
	return sign(l.apiSecret, nonce, timestampMsec, method, path, query)
}

func (l LBD) RetrieveServiceInformation(serviceId string) (string, error) {
	path := "/v1/services/" + serviceId

	ret, err := l.get(path, "", true)
	if err != nil {
		return "", err
	}

	fmt.Println(ret)
	return "ok", nil
}

func (l LBD) RetrieveServerTime() (int64, error) {
	path := "/v1/time"
	ret, err := l.get(path, "", false)
	if err != nil {
		return 0, err
	}
	return ret.ResponseTime, nil
}

func (l LBD) get(path, query string, sign bool) (*Response, error) {
	url := l.baseURL + path

	fmt.Println(url)

	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("service-api-key", l.apiKey)

	if sign {
		timeMsec := NowMsec()
		nonce := GenerateNonce(timeMsec)
		sig := l.Sign(nonce, timeMsec, "GET", path, query)

		fmt.Println(nonce, fmt.Sprint(timeMsec), sig)
		req.Header.Add("nonce", nonce)
		req.Header.Add("timestamp", fmt.Sprint(timeMsec))
		req.Header.Add("signature", sig)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := new(Response)
	err = json.Unmarshal(body, ret)
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

func sign(secret, nonce string, timestampMsec int64, method, path string, query string) string {
	var msg string
	if query == "" {
		msg = fmt.Sprintf("%s%d%s%s", nonce, timestampMsec, method, path)
	} else {
		msg = fmt.Sprintf("%s%d%s%s?%s", nonce, timestampMsec, method, path, query)
	}

	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(msg))
	sig := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(sig)
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
