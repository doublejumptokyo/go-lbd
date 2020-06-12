package lbd

import (
	"math/big"
	"testing"
)

var serviceTokenContractId = "d5e19f47"

func TestListAllServiceTokens(t *testing.T) {
	assert := initializeTest(t)
	ret, err := l.ListAllServiceTokens()
	assert.Nil(err)
	t.Log(ret[0])
}

func TestRetrieveServiceTokenInformation(t *testing.T) {
	assert := initializeTest(t)
	ret, err := l.RetrieveServiceTokenInformation(serviceTokenContractId)
	assert.Nil(err)
	t.Log(ret)
}

func TestMintServiceToken(t *testing.T) {
	onlyTxMode(t)
	assert := initializeTest(t)
	ret, err := l.MintServiceToken(serviceTokenContractId, toAddress, big.NewInt(1000))
	assert.Nil(err)
	t.Log(ret)
}
