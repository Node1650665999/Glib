package io

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// ParseOsArgs 解析 os.Args
func ParseOsArgs(filename *string, args *[]string) bool {
	if len(os.Args) > 0 {
		//os.Args[0] 是文件名, os.Args[1-n] 是用户参数
		for index, arg := range os.Args {
			if index == 0 {
				*filename = arg
			} else {
				*args = append(*args, arg)
			}
		}
		return true
	}
	return false
}

func ParseFlag(args [][]interface{}) {
	data := map[string]interface{}{}

	for _, arg := range args {
		parseVar(arg, data)
	}

	flag.Parse()

	for key, val := range data {
		switch val.(type) {
		case *string:
			data[key] = *(val.(*string))
		case *int:
			data[key] = *(val.(*int))
		case *bool:
			data[key] = *(val.(*bool))
		}
	}

	fmt.Println(data)
}

func parseVar(arg []interface{}, data map[string]interface{}) error {
	switch arg[1].(type) {
	case string:
		tmp := ""
		flag.StringVar(&tmp, arg[0].(string), arg[1].(string), arg[2].(string))
		data["name"] = &tmp
		return nil
	case int:
		tmp := 0
		flag.IntVar(&tmp, arg[0].(string), arg[1].(int), arg[2].(string))
		data["age"] = &tmp
	case bool:
		tmp := false
		flag.BoolVar(&tmp, arg[0].(string), arg[1].(bool), arg[2].(string))
		data["married"] = &tmp
	default:
		return fmt.Errorf("no match case")
	}
	return nil
}


