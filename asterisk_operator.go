package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{
		City:     "singkawang",
		Province: "kalbar",
		Country:  "Indo",
	}

	var address2 *Address = &address1

	//hen update address 2 to new pointer

	fmt.Println(address1)
	fmt.Println(address2)

	//when we change reference address2 to new pointer, all old reference on address1 is still not change
	address2 = &Address{
		City:     "sambas",
		Province: "kalbar",
		Country:  "Indo",
	}
	fmt.Println(address2)

	//we can add asterisk operator to make it two data binding, so when we change reference address2, all same reference with address2 will
	//refer to new address
	*address2 = Address{
		City:     "Jakarta",
		Province: "kalbar",
		Country:  "Indo",
	}
	fmt.Println(address2)
}
