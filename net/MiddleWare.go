package net

import (
	"fmt"
	"net/http"
)

//中间件处理函数
type MiddleWareHandle func(rw http.ResponseWriter, req *http.Request) error

//中间件结构体
type MiddleWare struct {
	//用来存储需要拦截的URI
	filterMap map[string]MiddleWareHandle
}

//初始化中间件
func NewMiddleWare() *MiddleWare {
	return &MiddleWare{filterMap: make(map[string]MiddleWareHandle)}
}

//注册中间件函数
func (f *MiddleWare) RegisterMiddleWareHandle(mw string, handler MiddleWareHandle) {
	f.filterMap[mw] = handler
}

//根据Uri获取对应的handle
func (f *MiddleWare) GetMiddleWareHandle(mw string) MiddleWareHandle {
	return f.filterMap[mw]
}

//声明 HttpHandle，用作处理请求的 http.HandleFunc 的参数
type HttpHandle func(w http.ResponseWriter, r *http.Request)

//运行中间件函数，返回 HttpHandle
func (f *MiddleWare) Handle(next HttpHandle, mw []string) HttpHandle {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range mw {
			handle,ok := f.filterMap[m]
			fmt.Printf("%v,%T",m, f.filterMap[m])
			if ! ok {
				continue
			}

			err := handle(w, r)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}

		//业务逻辑
		next(w, r)
	}
}