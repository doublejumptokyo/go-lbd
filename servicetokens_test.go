package lbd

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var serviceTokenContractId = "d5e19f47"
var name = "test_user"
var meta = "Test Api Function"

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

func TestUpdateServiceTokenInformation(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.UpdateServiceTokenInformation(serviceTokenContractId, name, meta)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestBurnServiceToken(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.BurnServiceToken(serviceTokenContractId, toAddress, big.NewInt(1000))
	assert.Nil(t, err)
	t.Log(ret)
}

func TestListAllServiceTokenHolders(t *testing.T) {
	onlyTxMode(t)
	ret, err := l.ListAllServiceTokenHolders(serviceTokenContractId)
	assert.Nil(t, err)
	t.Log(ret[0])
}
