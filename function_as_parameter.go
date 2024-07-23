package main

import "fmt"

type Filter func(string) string

func helloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func filterBadWords(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	helloWithFilter("anjing", filterBadWords)
}
