package main

import "fmt"

func applicationRunning() {
	defer logging()
	fmt.Println("application running")
}

func logging() {
	fmt.Println("logging")
}
func main() {
	applicationRunning()
}
