package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func ChangeCountryToIndonesiaWithPointer(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	address := Address{}
	address2 := new(Address)
	ChangeCountryToIndonesia(address)
	ChangeCountryToIndonesiaWithPointer(address2)

	fmt.Println(address)
	fmt.Println(address2)
}
