package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func PassByValue() {
	//pass by value
	address1 := Address{
		City:     "New York",
		Province: "New York",
		Country:  "USA",
	}

	address2 := address1

	address2.City = "Selakau" //whenc change address 2, address 1 does not affect

	fmt.Println(address1)
	fmt.Println(address2)
}

func PassByReference() {
	//pass by reference
	var address3 Address = Address{
		City:     "New York",
		Province: "New York",
		Country:  "USA",
	}

	var address4 *Address = &address3
	fmt.Println(address3)
	fmt.Println(address4)
}

func main() {
	PassByValue()
	PassByReference()
}
