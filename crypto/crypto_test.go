package crypto_test

import (
	"Glib/crypto"
	"testing"
)

func TestAES(t *testing.T) {
	name := "tcl"
	encodeStr := crypto.EncodeByAES(name)
	decodeStr := crypto.DecodeByAES(encodeStr)
	t.Log(encodeStr, decodeStr)
}