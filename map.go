package main

import "fmt"

func main() {
	var child map[string]string = map[string]string{}
	person := map[string]string{
		"name": "iqbal",
		"age":  "25",
	}

	child["name"] = "haha"
	child["age"] = "25"

	fmt.Println(person)
	fmt.Println(child)

	book := make(map[string]string)
	book["title"] = "Little Prince"
	book["author"] = "antoine"
	book["ups"] = "this is wrong"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
