package main

import "fmt"

func getFullName() (string, string) {
	return "iqbal", "atma"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	//ignore

	firstName, _ = getFullName()
	fmt.Println(firstName)
}
