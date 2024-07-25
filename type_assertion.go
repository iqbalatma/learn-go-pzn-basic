package main

import "fmt"

func Random() interface{} {
	return "OK"
}

func main() {
	result := Random()
	resultString := result.(string)
	fmt.Println(resultString)

	//type assertion with switch

	result = Random()
	switch value := result.(type) {
	case string:
		fmt.Println("ini tuh ", value)
	case int:
		fmt.Println("ini tuh ", value)
	}
}
