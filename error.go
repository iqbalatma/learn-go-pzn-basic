package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := Pembagian(100, 10)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
