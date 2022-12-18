package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2

	fmt.Println(myMap["First"], myMap["Second"])

}
