package main

import "Glib/io"

func main() {



	/*flagArgs := [][]interface{}{
		{"name", "tcl", "姓名"},
		{"age", 18, "年龄"},
		{"married", false, "婚否"},
	}*/

	flagArgs := []io.Arg{
		{Option: "name", Default: "tcl", Remark: "姓名"},
		{Option: "age", Default:  18, Remark: "年龄"},
		{Option: "married", Default: false, Remark: "婚否"},
	}

	io.ParseFlag(flagArgs)

}
