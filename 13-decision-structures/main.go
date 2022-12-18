package main

import "fmt"

func main() {
	cat := "cat"

	if cat == "cat" {
		fmt.Println("Cat is equal to", cat)
	} else {
		fmt.Println("Cat is not cat, it is", cat)
	}
}
