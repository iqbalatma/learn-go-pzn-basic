package main

import "fmt"

func endApp() {
	fmt.Println("end app")
	message := recover()
	fmt.Println(message)
}

func errorCheck(isError bool) {
	defer endApp()
	if isError {
		panic("ERROR POKOKNYA")
	}
}
func main() {

	errorCheck(true)
}
