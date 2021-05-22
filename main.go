package main

import (
	"Glib/net"
	"encoding/json"
	"fmt"

	/*"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"Glib/net"*/
	"os"
	"path/filepath"
)

type TvInfo struct {
	Id            string   `json:"id"`
	ClassType     string   `json:"class_type"`
	DirectoryType string   `json:"directory_type"`
	Url           string   `json:"url"`
	CoverUrl      string   `json:"cover_url"`
	Title         string   `json:"title"`
	DirectoryName string   `json:"directory_name"`
	SetName       string   `json:"set_name"`
	Words         []string `json:"words"`
}

type Data struct {
	List []TvInfo `json:"list"`
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int    `json:"time"`
	Data Data   `json:"data"`
}

func main() {

	params := map[string]string{
		"platform" : "1",
		"token"    : "a9217f0271c2b70ea2ee327eec47d58d",
		"tv_type"  : "1",
	}

	 var res  Result
	data,_ := net.Get("http://test-newapi.hanzigon.cn/pc/v2/tv/class_info", params)

	json.Unmarshal([]byte(data), &res)

	fmt.Printf("%+v",res)

}

func Filepath(params ...string) string {
	wd, _ := os.Getwd()
	params = append([]string{wd}, params...)
	//dirSep := string(os.PathSeparator)
	return filepath.Join(params...)
}


