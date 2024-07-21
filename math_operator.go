package main

import "fmt"

func main() {
	var a int8 = 10
	var b int8 = 10
	var c int8 = a + b
	var d int8 = 5
	var e int8 = 2
	fmt.Println(c)
	fmt.Println(d * e)

	//augmented assignment
	a = a + 1
	a += 1
	fmt.Println(a)

	//unary operator
	var j = 1
	j++
	j++
	fmt.Println(j)

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a != b)
}
