package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Printf("Hello %s, my name is %s", name, customer.Name)
}

func main() {
	var person Customer
	person.Name = "Iqbal"
	person.Address = "Singkawang"
	person.Age = 25

	fmt.Println(person)

	var person2 = Customer{
		Name:    "Iqbal",
		Address: "Singkawang",
		Age:     25,
	}

	fmt.Println(person2)

	var person3 = Customer{"iqbal", "skw", 20}

	fmt.Println(person3)
	person3.sayHello("joko")
}
