package main

import (
	"Glib/io"
	"fmt"
)

func main() {
	name    := io.Arg{Option: "name", Default: "tcl", Remark: "姓名"}
	age     := io.Arg{Option: "age", Default:  18, Remark: "年龄"}
	married := io.Arg{Option: "married", Default: false, Remark: "婚否"}

	flagArgs := []*io.Arg{&name, &age, &married}

	io.ParseFlag(flagArgs)

	fmt.Printf("name:%+v \n age:%+v \n married:%+v", name, age, married)

}
