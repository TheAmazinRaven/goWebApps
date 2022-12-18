package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Sheba"
	myMap["other-dog"] = "Rusty"

	fmt.Println(myMap["dog"])
	fmt.Println(myMap["other-dog"])
}
