package main

import "fmt"

func main() {
	var counter int8 = 1
	for counter <= 10 {
		fmt.Println("this is counter number ", counter)
		counter++
	}

	//for with statement
	//init statement, condition check, post statement
	for i := 0; i < 10; i++ {
		fmt.Println("this is i", i)
	}

	//manual loop on map
	var numbers [3]int8 = [3]int8{1, 2, 3}
	var person map[string]string = map[string]string{
		"name": "iqbal",
		"age":  "25",
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println(person)

	names := []string{"iqbal", "atma", "muliawan"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
