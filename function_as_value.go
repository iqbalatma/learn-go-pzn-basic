package main

import "fmt"

func getGoodBye() string {
	return "GOOD BYE"
}

func main() {
	goodBye := getGoodBye

	fmt.Println(goodBye())
}
