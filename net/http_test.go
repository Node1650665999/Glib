package net_test

import (
	mynet "github.com/Node1650665999/Glib/net"
	"net/http"
	"testing"
)

func TestApiResponse(t *testing.T) {
	// 处理请求
	http.HandleFunc("/remote-ip", HttpHandler)
	http.ListenAndServe(":8084", nil)
}
func HttpHandler(w http.ResponseWriter, r *http.Request)  {
	remoteIp := mynet.RemoteIp(r)
	data := map[string]interface{}{"ip": remoteIp}
	mynet.ApiResponse(200, "成功", data, w)
	return

}

func TestProxyForward(t *testing.T) {
	http.HandleFunc("/proxy", HandleProxy)
	http.ListenAndServe(":8085", nil)
}
func HandleProxy(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("success"))
	//mynet.ProxyForward(w, r, "http://www.baidu.com:80/")
}