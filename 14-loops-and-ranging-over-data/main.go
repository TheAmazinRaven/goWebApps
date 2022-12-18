package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func main() {

	var users []User
	users = append(users, User{"Rae", "D.", "email@awesome.com", 29})
	users = append(users, User{"Sally", "M.", "email@awesome.com", 41})
	users = append(users, User{"Bob", "S.", "email@awesome.com", 12})
	users = append(users, User{"Tenoch", "H.", "email@awesome.com", 35})

	for _, l := range users {
		fmt.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
