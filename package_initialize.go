package main

import (
	"fmt"
	"learn-go-pzn-basic/database"
	_ "learn-go-pzn-basic/internal"
)

func main() {
	fmt.Println(database.GetConnection())
}
