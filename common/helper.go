package common

import (
	"bytes"
	"fmt"
	"strings"
)

//InArray 判断切片是否包含某个元素
func InArray(need interface{}, haystack interface{}) bool {
	switch key := need.(type) {
	case int:
		for _, item := range haystack.([]int) {
			if item == key {
				return true
			}
		}
	case string:
		for _, item := range haystack.([]string) {
			if item == key {
				return true
			}
		}
	case int64:
		for _, item := range haystack.([]int64) {
			if item == key {
				return true
			}
		}
	case float64:
		for _, item := range haystack.([]float64) {
			if item == key {
				return true
			}
		}
	default:
		return false
	}
	return false
}

//ByteFormat 将字节格式化为指定的单位
// refer:https://blog.microdba.com/golang/2016/05/01/golang-byte-conv/
func ByteFormat(size uint64) string {
	sz   := float64(size)
	base := float64(1024)
	unit := []string{"B","KB","MB","GB","TB","EB"}
	i := 0

	for sz >= base {
		sz /= base
		i++
	}
	return fmt.Sprintf("%.2f%s",sz, unit[i])
}


// Ext 获取文件扩展名
func Ext(file string) (basename string, ext string) {
	arr := strings.Split(file, ".")
	return arr[0], arr[len(arr)-1]
}

//StrJoin 用来拼接字符串
func StrJoin(args ...string) string {
	var buf bytes.Buffer
	for _, arg := range args {
		buf.WriteString(arg)
	}
	return buf.String()
}