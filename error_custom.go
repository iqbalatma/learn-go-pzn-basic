package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}

func SavedData(id string, data any) error {
	if id == "" {
		return &ValidationError{"id cannot be empty"}
	}

	if id != "iqbal" {
		return &NotFoundError{"Id not found"}
	}

	return nil
}

func main() {
	err := SavedData("", nil)

	fmt.Println(err)
}
