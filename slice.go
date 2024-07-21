package main

import "fmt"

func main() {
	names := [...]string{"iqbal", "atma", "muliawan", "jos", "joko", "mantap"}
	var slice []string = names[4:]
	fmt.Println(cap(slice))
	fmt.Println(slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))

	fromSlice := names[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(toSlice)
}
