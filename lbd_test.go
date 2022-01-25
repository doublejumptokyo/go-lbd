package lbd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	l                      = &LBD{}
	serviceID              = os.Getenv("SERVICE_ID")
	owner                  = NewWallet(os.Getenv("OWNER_ADDR"), os.Getenv("OWNER_SECRET"))
	itemTokenContractId    = os.Getenv("ITEMTOKEN_CONTRACT_ID")
	serviceTokenContractId = os.Getenv("SERVICETOKEN_CONTRACT_ID")
	tokenType              = "10000001"
	fungibleTokenType      = "00000001"
	userId                 = os.Getenv("USER_ID")
	toAddress              = userId
	sessionToken           = os.Getenv("SESSION")
)

func TestSign(t *testing.T) {
	assert := assert.New(t)
	// https://docs-blockchain.line.biz/api-guide/Authentication
	key := "136db0ad-0fe1-456f-96a4-329be3f93036"
	secret := "9256bf8a-2b86-42fe-b3e0-d3079d0141fe"
	nonce := "Bp0IqgXE"
	timestamp := int64(1581850266351)

	l, err := NewCashew(key, secret, nil)
	assert.Nil(err)

	// Example 1
	ex1 := &Request{nonce, timestamp, "GET", "/v1/wallets", nil}
	sig1 := l.Sign(ex1)
	expected1 := "2LtyRNI16y/5/RdoTB65sfLkO0OSJ4pCuz2+ar0npkRbk1/dqq1fbt1FZo7fueQl1umKWWlBGu/53KD2cptcCA=="
	assert.Equal(sig1, expected1)

	// Example 3
	ex2 := &UpdateNonFungibleInformationRequest{
		Request:      &Request{nonce, timestamp, "PUT", "/v1/item-tokens/61e14383/non-fungibles/10000001/00000001", nil},
		OwnerAddress: "tlink1fr9mpexk5yq3hu6jc0npajfsa0x7tl427fuveq",
		OwnerSecret:  "uhbdnNvIqQFnnIFDDG8EuVxtqkwsLtDR/owKInQIYmo=",
		Name:         "NewName",
	}
	t.Log(ex2.Encode())
	sig2 := l.Sign(ex2)
	t.Log(sig2)
	expected2 := "4L5BU0Ml/ejhzTg6Du12BDdElv8zoE7XD/iyOaZ2BHJIJG0SUOuCZWXu0YaF4i4C2CFJhjZoJFsje4CJn/wyyw=="
	assert.Equal(sig2, expected2)
}

func TestMain(m *testing.M) {
	var err error
	l, err = NewCashew(os.Getenv("API_KEY"), os.Getenv("API_SECRET"), owner)
	if err != nil {
		panic(err)
	}

	if os.Getenv("DEBUG") != "" {
		l.Debug = true
	}

	status := m.Run()
	os.Exit(status)
}

func onlyTxMode(t *testing.T) {
	t.Helper()
	if !(os.Getenv("TX") == "1") {
		t.Skip("skipping test in no Tx mode. Set env TX=1")
	}
}
