package main

import "Glib/io"

func main() {

	/*var name string
	var age  int
	var status bool*/

	flagArgs := [][]interface{}{
		{"name", "tcl", "姓名"},
		{"age", 18, "年龄"},
		{"married", false, "婚否"},
	}

	//io.ParseFlag(name, age, status)
	io.ParseFlag(flagArgs)

	//io.Test()
}
