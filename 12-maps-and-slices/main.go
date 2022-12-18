package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 9)
	mySlice = append(mySlice, 7)
	mySlice = append(mySlice, 24)

	fmt.Println(mySlice)

	sort.Ints(mySlice)

	fmt.Println(mySlice)
}
