package net

import (
	"io/ioutil"
	"net/http"
	"net/url"
)



//Get 发送get请求
func Get(apiUrl string, param map[string]string) (string,error) {
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
