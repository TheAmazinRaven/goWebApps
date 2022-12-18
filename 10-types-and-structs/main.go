package main

import (
	"fmt"
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var s2 = "six"

	s := "eight"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

	userStruct()
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the say something func is", s)
	return s3, "world"
}

func userStruct() {
	user := User{
		FirstName:   "Rae",
		LastName:    "D.",
		PhoneNumber: "1 555 555-1212",
	}

	fmt.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate, user.PhoneNumber)
}
