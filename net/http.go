package net

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//GetSimple 发送普通get请求
func GetSimple(apiUrl string, param map[string]string) (string,error) {
	data := url.Values{}
	for name, val := range param {
		data.Set(name, val)
	}
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return "", err
	}

	u.RawQuery = data.Encode() // URL encode
	//fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	return string(b), err
}

// Get 发送复杂 get 请求
func Get(apiUrl string, param map[string]string, header map[string]string) (string,error) {
	// 创建请求
	req, _ := http.NewRequest("GET", apiUrl, nil)

	// 设置参数
	q := req.URL.Query()
	for name, val := range param {
		q.Add(name, val)
	}
	req.URL.RawQuery = q.Encode()

	// 设置请求头
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// 创建 Http 客户端
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// 关闭连接
	defer resp.Body.Close()

	//读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// PostSimple 发送简单的 post 请求(默认 Content-Type=multipart/form-data)
func PostSimple(apiUrl string, param map[string]string) (string,error) {
	//处理参数
	data := url.Values{}
	for name, val := range param {
		data.Set(name, val)
	}

	resp, err := http.PostForm(apiUrl, data)
	if err != nil {
		return "", err
	}

	// 关闭客户端
	defer resp.Body.Close()

	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(body), nil
}

//Post 发送复杂 post 请求
func Post(apiUrl string, param map[string]string, header map[string]string) (string,error)  {
	// 设置参数
	data := url.Values{}
	for name, val := range param {
		data.Set(name, val)
	}

	// 创建请求
	req, _ := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// 创建 Http 客户端
	client := &http.Client{}
	resp ,err := client.Do(req)
	if err != nil {
		return "", nil
	}

	// 关闭连接
	defer resp.Body.Close()

	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}

//PostJson 发送 Content-Type=application/json 的请求
func PostJson(apiUrl string, param interface{}, header map[string]string) (string,error) {
	// 转义参数
	paramJson,_ := json.Marshal(param)

	// 创建请求
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(paramJson))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	for k, v := range  header {
		req.Header.Set(k, v)
	}

	// 创建 Http 客户端
	client    := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// 关闭连接
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//GetParam 获取 Http Get请求参数
func GetParam(key string,req *http.Request) string {
	req.ParseForm()
	return req.FormValue(key)
}

//PostParam 获取 Http Post请求参数
func PostParam(key string,req *http.Request) string {
	req.ParseForm()
	return req.Form.Get(key)
}

//Cors 设置跨域
func Cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true") //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")             //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}