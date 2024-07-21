package main

import "fmt"

func main() {
	var name string = "iqbal"
	if name == "iqbal" {
		fmt.Println(name)
	} else if name == "joko" {
		fmt.Println("hi, joko")
	} else {
		fmt.Println("hi, can i get your number")
	}

	if length := len(name); length > 100 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("name oke")
	}
}
