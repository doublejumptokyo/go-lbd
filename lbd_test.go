package lbd

import (
	"os"
	"testing"

	"github.com/cheekybits/is"
)

var (
	l         = &LBD{}
	serviceID = os.Getenv("SERVICE_ID")
	owner     = &Account{os.Getenv("OWNER_ADDR"), os.Getenv("OWNER_SECRET")}
	toAddress = "tlink13j9ctt0r05q7hq7syf34qm973hl5hftk9m662g"
)

func TestSign(t *testing.T) {
	is := is.New(t)
	// https://docs-blockchain.line.biz/api-guide/Authentication
	key := "136db0ad-0fe1-456f-96a4-329be3f93036"
	secret := "9256bf8a-2b86-42fe-b3e0-d3079d0141fe"
	nonce := "Bp0IqgXE"
	timestamp := int64(1581850266351)

	l, err := NewLBD(key, secret)
	is.Nil(err)

	// Example 1
	ex1 := &Request{nonce, timestamp, "GET", "/v1/wallets"}
	sig1 := l.Sign(ex1)
	expected1 := "2LtyRNI16y/5/RdoTB65sfLkO0OSJ4pCuz2+ar0npkRbk1/dqq1fbt1FZo7fueQl1umKWWlBGu/53KD2cptcCA=="
	is.Equal(sig1, expected1)

	// Example 3
	ex2 := &MintNonFungibleRequest{
		Request:      &Request{nonce, timestamp, "PUT", "/v1/item-tokens/61e14383/non-fungibles/10000001/00000001"},
		OwnerAddress: "tlink1fr9mpexk5yq3hu6jc0npajfsa0x7tl427fuveq",
		OwnerSecret:  "uhbdnNvIqQFnnIFDDG8EuVxtqkwsLtDR/owKInQIYmo=",
		Name:         "NewName",
	}
	t.Log(ex2.Encode())
	sig2 := l.Sign(ex2)
	t.Log(sig2)
	expected2 := "4L5BU0Ml/ejhzTg6Du12BDdElv8zoE7XD/iyOaZ2BHJIJG0SUOuCZWXu0YaF4i4C2CFJhjZoJFsje4CJn/wyyw=="
	is.Equal(sig2, expected2)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	l, err = NewLBD(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	is.Nil(err)
	return is
}
