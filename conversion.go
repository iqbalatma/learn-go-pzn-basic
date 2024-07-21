package main

import "fmt"

func main() {
	var value32 int32 = 123
	var value64 int64 = int64(value32)
	fmt.Println(value64)

	var name string = "Iqbal atma muliawan"
	var e byte = name[3]
	var s string = string(e)
	fmt.Println(e)
	fmt.Println(s)
}
