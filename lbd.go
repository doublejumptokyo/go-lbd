package lbd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	CashewBaseURL = "https://test-api-blockchain.line.me"
)

type LBD struct {
	baseURL string
	APIKey  string
}

type Response struct {
	ResponseTime  int64       `json:"responseTime"`
	StatusCode    int64       `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	ResponseData  interface{} `json:"responseData"`
}

func NewLBD(apiKey string) (*LBD, error) {
	l := &LBD{
		baseURL: CashewBaseURL,
		APIKey:  apiKey,
	}
	return l, nil
}

func (l LBD) RetrieveServerTime() (int64, error) {
	path := "/v1/time"

	ret, err := l.get(path)
	if err != nil {
		return 0, err
	}

	fmt.Println(ret)
	return 0, nil
}

func (l LBD) get(path string) (*Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := l.baseURL + path

	fmt.Println(url)

	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("service-api-key", l.APIKey)

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
