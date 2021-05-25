package main

import (
	"Glib/net"
	"fmt"
	sj "github.com/bitly/go-simplejson"
)



func main() {
	params := map[string]interface{}{
		"platform" : "1",
		"token"    : "a9217f0271c2b70ea2ee327eec47d58d",
		"tv_type"  : "1",
	}

	response,_ := net.PostJson("http://test-newapi.hanzigon.cn/pc/v2/tv/class_info",params, nil)
	res,_      := sj.NewJson([]byte(response))
	data       := res.Get("data")
	fmt.Println(data)





}



