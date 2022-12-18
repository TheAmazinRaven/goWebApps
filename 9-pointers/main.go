package main

import "log"

func main() {
	var myString string
	myString = "Pink"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("After the function has been called, myString is set to", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Purple"
	*s = newValue
}
