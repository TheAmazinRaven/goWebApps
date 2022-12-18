package main

import "fmt"

func main() {
	myVar := "apple"

	switch myVar {
	case "cat":
		fmt.Println("Cat is set to cat.")
	case "dog":
		fmt.Println("Cat is set to dog.")
	case "fish":
		fmt.Println("Cat is set to fish.")

	default:
		fmt.Println("Cat is set to something else.")
	}

}
