package main

import (
	"fmt"
	"strconv"
)

func main() {
	var slice []byte = make([]byte,0,1024)
	slice = strconv.AppendBool(slice,false)
	slice = strconv.AppendInt(slice,1234,10)
	fmt.Println(string(slice))
}

func test02() {
	var str string
	str = strconv.FormatFloat(3.14156383838383, 'f', 3, 64)
	fmt.Println(str)
}
func test01()  {
	var str string
	str = strconv.FormatBool(false)
	fmt.Println(str)
}
