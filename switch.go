package main

import "fmt"

func main() {
	name := "iqbal"

	switch name {
	case "iqbal":
		fmt.Println("helo iqbal")
	case "budi":
		fmt.Println("helo budi")
	default:
		fmt.Println("hi what is your name")
	}

}
