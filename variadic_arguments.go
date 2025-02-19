package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(1, 2, 3)

	fmt.Println(result)

	//we also can destruct slice to pass as varargs
	numbers := []int{1, 2, 3, 4, 5}
	result2 := sumAll(numbers...)
	fmt.Println(result2)
}
