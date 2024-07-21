package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "iqbal"
	names[1] = "atma"
	names[2] = "muliawan"
	fmt.Println(names)

	var values = [3]int{
		1, 2, 3,
	}
	fmt.Println(values)
}
