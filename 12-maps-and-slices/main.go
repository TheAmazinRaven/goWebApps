package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Rae",
		LastName:  "D.",
	}

	myMap["me"] = me

	fmt.Println(myMap["me"].FirstName, myMap["me"].LastName)

	var myNewVar float32
	myNewVar = 11.1

	fmt.Println(myNewVar)

}
