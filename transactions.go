package lbd

import (
	"encoding/json"
)

func UnmarshalTransactionInformation(data []byte) (TransactionInformation, error) {
	var r TransactionInformation
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TransactionInformation) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TransactionInformation struct {
	Height    int64  `json:"height"`
	Txhash    string `json:"txhash"`
	Index     int64  `json:"index"`
	Code      int64  `json:"code"`
	Logs      []Log  `json:"logs"`
	GasWanted int64  `json:"gasWanted"`
	GasUsed   int64  `json:"gasUsed"`
	Tx        Tx     `json:"tx"`
	Timestamp string `json:"timestamp"`
}

type Log struct {
	MsgIndex int64   `json:"msgIndex"`
	Success  bool    `json:"success"`
	Log      string  `json:"log"`
	Events   []Event `json:"events"`
}

type Event struct {
	Type       string      `json:"type"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tx struct {
	Type  string  `json:"type"`
	Value TxValue `json:"value"`
}

type TxValue struct {
	Msg        []Msg       `json:"msg"`
	Fee        Fee         `json:"fee"`
	Memo       string      `json:"memo"`
	Signatures []Signature `json:"signatures"`
}

type Fee struct {
	Gas    int64         `json:"gas"`
	Amount []interface{} `json:"amount"`
}

type Msg struct {
	Type  string   `json:"type"`
	Value MsgValue `json:"value"`
}

type MsgValue struct {
	From       string `json:"from"`
	ContractID string `json:"contractId"`
	To         string `json:"to"`
	Amount     int64  `json:"amount"`
}

type Signature struct {
	PubKey    PubKey `json:"pubKey"`
	Signature string `json:"signature"`
}

type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (l LBD) RetrieveTransactionInformation(txHash string) (*TransactionInformation, error) {
	r := NewGetRequest("/v1/transactions/" + txHash)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	ret := new(TransactionInformation)
	return ret, json.Unmarshal(resp.ResponseData, ret)
}
