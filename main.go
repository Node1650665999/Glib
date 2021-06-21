package main

import (
	"Glib/common"
	"fmt"
)

func main()  {
	file := "text.img"
	ext,_ := common.Ext(file)
	fmt.Println(ext)
}

