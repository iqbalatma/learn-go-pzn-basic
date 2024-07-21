package main

import "fmt"

func main() {
	names := []string{"iqbal", "atma", "muliawan"}

	for i, name := range names {
		if name == "atma" {
			continue
		}
		fmt.Println(i, name)
	}
}
