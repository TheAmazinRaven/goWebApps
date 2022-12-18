package main

import "fmt"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Rae"

	myVar2 := myStruct{
		FirstName: "Bonita",
	}

	fmt.Println("My var is set to", myVar, "& myVar2 is set to", myVar2)

	fmt.Println("myVar is set to", myVar.printFirstName())
	fmt.Println("myVar2 is set to", myVar2.printFirstName())
}
