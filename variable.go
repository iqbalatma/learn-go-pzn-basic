package main

import "fmt"

func main() {
	var name string
	name = "iqbal atma muliawan"
	fmt.Println("Enter your name:", name)

	name = "namanya berubah"
	fmt.Println("Name changes:", name)

	var nameWithoutDataType = "iqbal"
	fmt.Println("Enter your name:", nameWithoutDataType)

	nameWithoutVar := "nama tanpa var"
	fmt.Println("Enter your name:", nameWithoutVar)

	nameWithoutVar = "variable changes without : because data change, not declaration"

	var (
		firstName  = "iqbal"
		middleName = "atma"
		lastName   = "muliawan"
	)

	fmt.Println(firstName, middleName, lastName)
}
