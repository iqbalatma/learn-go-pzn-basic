package main

import "fmt"

func getFactorial(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func getFactorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * getFactorialRecursive(value-1)
	}
}
func main() {
	fmt.Println(getFactorial(3))
	fmt.Println(getFactorialRecursive(3))
}
