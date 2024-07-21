package main

import "fmt"

func main() {
	names := []string{"iqbal", "atma", "muliawan"}

	for i, name := range names {
		if name == "muliawan" {
			break
		}
		fmt.Println(i, name)
	}
}
