package lbd

import (
	"testing"
)

func TestCreateNonFungible(t *testing.T) {
	is := initializeTest(t)
	ret, err := l.CreateNonFungible(serviceID, "NobunagaOda", "Tenkafubu", owner)
	is.Nil(err)
	t.Log(ret)
}
