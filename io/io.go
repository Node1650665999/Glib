package io

import (
	"flag"
	"fmt"
	"os"
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

type Arg struct {
	Option string
	Default interface{}
	Remark  string
	Value  interface{}
}

func ParseFlag(args []Arg) {
	data := map[string]interface{}{}

	for _, arg := range args {
		setVar(arg, data)
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

// setVar 将选项和值绑定到map中
func setVar(arg Arg, data map[string]interface{}) error {

	option := arg.Option
	remark := arg.Remark

	switch arg.Default.(type) {
	case string:
		tmp := ""
		flag.StringVar(&tmp, option, arg.Default.(string), remark)
		data[option] = &tmp
		return nil
	case int:
		tmp := 0
		flag.IntVar(&tmp, option, arg.Default.(int), remark)
		data[option]  = &tmp
	case bool:
		tmp := false
		flag.BoolVar(&tmp, option, arg.Default.(bool), remark)
		data[option] = &tmp
	default:
		return fmt.Errorf("no match case")
	}
	return nil
}


