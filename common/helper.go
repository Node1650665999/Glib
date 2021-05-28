package common

import "strings"

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

// Ext 获取文件扩展名
func Ext(file string) (basename string, ext string) {
	arr := strings.Split(file, ".")
	return arr[0], arr[len(arr)-1]
}
