package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "iqbal"
	lastName = "atma"
	return firstName, lastName
}

func main() {
	firstName, _ := getCompleteName()
	fmt.Println(firstName)
}
