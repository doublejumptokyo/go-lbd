package lbd

import (
	"math/big"
	"testing"
)

var serviceTokenContractId = "d5e19f47"

func TestListAllServiceTokens(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.ListAllServiceTokens()
	is.Nil(err)
	t.Log(ret[0])
}

func TestRetrieveServiceTokenInformation(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.RetrieveServiceTokenInformation(serviceTokenContractId)
	is.Nil(err)
	t.Log(ret)
}

func TestMintServiceToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	is := initializeTest(t)
	ret, err := l.MintServiceToken(serviceTokenContractId, toAddress, big.NewInt(1000), owner)
	is.Nil(err)
	t.Log(ret)
}
