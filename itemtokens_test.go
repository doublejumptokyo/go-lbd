package lbd

import (
	"testing"
)

var (
	itemTokenContractId = "a5e30e57"
)

func TestCreateNonFungible(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.CreateNonFungible(itemTokenContractId, "NobunagaOda", "Tenkafubu", owner)
	is.Nil(err)
	t.Log(ret)
}

func TestMintNonFungible(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.MintNonFungible(itemTokenContractId, "10000002", "Nobnyaga", "uwawa", toAddress, owner)
	is.Nil(err)
	t.Log(ret)
}
