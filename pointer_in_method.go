package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr " + man.Name
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr " + man.Name
}

func main() {
	iqbal := Man{"Iqbal"}

	iqbal.Married()
	iqbal.MarriedPointer()
	fmt.Println(iqbal)
}
