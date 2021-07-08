package crypto_test

import (
	"github.com/node1650665999/Glib/crypto"
	"testing"
)

func TestAES(t *testing.T) {
	name := "tcl"
	encodeStr := crypto.EncodeByAES(name)
	decodeStr := crypto.DecodeByAES(encodeStr)
	t.Log(encodeStr, decodeStr)
}