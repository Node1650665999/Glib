package net_test

import (
	. "github.com/node1650665999/Glib/net"
	"errors"
	"net/http"
	"testing"
)

func TestMiddleWare(t *testing.T) {
	// 实例化中间对象
	middleware := NewMiddleWare()
	// 注册中间件函数
	middleware.RegisterMiddleWareHandle("loginAuth", LoginAuth)

	// 处理请求
	http.HandleFunc("/user-info", middleware.Handle(UserInfo, []string{"loginAuth"}))
	http.ListenAndServe(":8083", nil)
}

//中间件函数
func LoginAuth(w http.ResponseWriter, r *http.Request) error {
	//http://localhost:8083/user-info?uid=1
	if GetParam("uid", r) == "1" {
		return errors.New("验证不通过")
	}
	return nil
}

//业务函数
func UserInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("我的姓名是xxx,年龄是xxx"))
}