package main

import "fmt"

func main() {
	type NoKTP string
	var ktpIqbal NoKTP = "6123"

	var contohKTPIqbal string = "1234"

	fmt.Println(ktpIqbal)
	fmt.Println(NoKTP(contohKTPIqbal))
}
