package main

import (
	"Glib/crypto"
	"fmt"
)
func main() {
	//str := "passwd"
	url := "www.baidu.com"

	fmt.Println(crypto.UrlEncode(url), crypto.EncodeByBase64(url))

}