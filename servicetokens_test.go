package lbd

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var serviceTokenContractId = "d5e19f47"

func TestListAllServiceTokens(t *testing.T) {
	ret, err := l.ListAllServiceTokens()
	assert.Nil(t, err)
	t.Log(ret[0])
}

func TestRetrieveServiceTokenInformation(t *testing.T) {
	ret, err := l.RetrieveServiceTokenInformation(serviceTokenContractId)
	assert.Nil(t, err)
	t.Log(ret)
}

func TestMintServiceToken(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.MintServiceToken(serviceTokenContractId, toAddress, big.NewInt(1000))
	assert.Nil(t, err)
	t.Log(ret)
}
