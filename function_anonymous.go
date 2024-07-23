package main

import "fmt"

type Blacklist func(string) bool

func blockedUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You are blacklisted", name)
	} else {
		fmt.Println("You are not blacklisted", name)
	}
}

func main() {
	blockedUser("iqbal", func(name string) bool {
		return name == "iqbal"
	})
}
