package main

import "fmt"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Rae")
	mySlice = append(mySlice, "Janet")
	mySlice = append(mySlice, "Henry")
	mySlice = append(mySlice, "Hannah")
	mySlice = append(mySlice, "MawMaw")

	fmt.Println(mySlice)
}
