package main

import (
	"Glib/crypto"
	"fmt"
)

func main () {
	name := "tcl"
	encodeStr := crypto.EncodeByAES(name)
	decodeStr := crypto.DecodeByAES(encodeStr)

	fmt.Println(encodeStr, decodeStr)
}
