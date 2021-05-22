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

func ParseFlag(args []*Arg) error{
	for _, arg := range args {
		err := setVar(arg)
		if err != nil {
			return err
		}
	}

	flag.Parse()

	for _, arg := range args {
		switch arg.Value.(type) {
		case *string:
			arg.Value = *(arg.Value.(*string))
		case *int:
			arg.Value = *(arg.Value.(*int))
		case *bool:
			arg.Value = *(arg.Value.(*bool))
		}
	}

	return nil
}

// setVar 将选项和值绑定到map中
func setVar(arg *Arg) error {
	option := arg.Option
	remark := arg.Remark

	switch arg.Default.(type) {
	case string:
		tmp := ""
		flag.StringVar(&tmp, option, arg.Default.(string), remark)
		arg.Value = &tmp
	case int:
		tmp := 0
		flag.IntVar(&tmp, option, arg.Default.(int), remark)
		arg.Value = &tmp
	case bool:
		tmp := false
		flag.BoolVar(&tmp, option, arg.Default.(bool), remark)
		arg.Value = &tmp
	default:
		return  fmt.Errorf("no match anything")
	}
	return nil
}


