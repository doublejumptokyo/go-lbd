package lbd

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
)

func UnmarshalTransaction(data []byte) (*Transaction, error) {
	r := new(Transaction)
	return r, json.Unmarshal(data, r)
}

func (r *Transaction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Transaction struct {
	Height    int64  `json:"height"`
	Txhash    string `json:"txhash"`
	Index     int64  `json:"index"`
	Codespace string `json:"codespace"`
	Code      int64  `json:"code"`
	Logs      []Log  `json:"logs"`
	GasWanted int64  `json:"gasWanted"`
	GasUsed   int64  `json:"gasUsed"`
	Tx        Tx     `json:"tx"`
	Timestamp int64  `json:"timestamp"`
}

func (t *Transaction) Check() (err error) {
	if t.Code == 0 {
		return nil
	}
	return fmt.Errorf("Transaction failure %s %d", t.Codespace, t.Code)
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
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type Signature struct {
	PubKey    PubKey `json:"pubKey"`
	Signature string `json:"signature"`
}

type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (l LBD) RetrieveTransactionInformation(txHash string) (*Transaction, error) {
	r := NewGetRequest("/v1/transactions/" + txHash)
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransaction(resp.ResponseData)
}

func (l LBD) GetExplorerURL(tx *Transaction) string {
	return fmt.Sprintf("https://explorer.blockchain.line.me/%s/transaction/%s", strings.ToLower(string(l.Network)), tx.Txhash)
}

type TransactionV2SummaryResult struct {
	Code      int64  `json:"code"`
	CodeSpace string `json:"codeSpace"`
	Status    string `json:"status"`
}

type TransactionV2Summary struct {
	Height    int64                       `json:"height"`
	TxIndex   int64                       `json:"txIndex"`
	TxHash    string                      `json:"txHash"`
	Timestamp int64                       `json:"timestamp"`
	Signer    []string                    `json:"signer"`
	Result    *TransactionV2SummaryResult `json:"result"`
}

type TransactionV2Message struct {
	MsgIndex    int64           `json:"msgIndex"`
	RequestType string          `json:"resultType"`
	Details     json.RawMessage `json:"details"`
}

type TransactionV2 struct {
	Summary  *TransactionV2Summary   `json:"summary"`
	Messages []*TransactionV2Message `json:"messages"`
	Events   []json.RawMessage       `json:"events"`
}

func (t *TransactionV2) Check() error {
	var (
		code      int64 = -1
		codespace string
	)
	if t.Summary != nil && t.Summary.Result != nil {
		code = t.Summary.Result.Code
		codespace = t.Summary.Result.CodeSpace
	}
	if code == 0 {
		return nil
	}
	return fmt.Errorf("Transaction failure %s %d", codespace, code)
}

func UnmarshalTransactionV2(data []byte) (*TransactionV2, error) {
	r := new(TransactionV2)
	return r, json.Unmarshal(data, r)
}

func (l LBD) RetrieveTransactionInformationV2(txHash string) (*TransactionV2, error) {
	r := NewGetRequest(path.Join("/v2/transactions", txHash))
	resp, err := l.Do(r, true)
	if err != nil {
		return nil, err
	}
	return UnmarshalTransactionV2(resp.ResponseData)
}
