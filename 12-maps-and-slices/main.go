package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numbers)

	fmt.Println(numbers[0:2])
	fmt.Println(numbers[6:9])
}
