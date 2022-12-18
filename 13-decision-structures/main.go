package main

import "fmt"

func main() {
	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		fmt.Println("myNum is greater than 99 & isTrue is set to true.")
	} else if myNum < 100 && isTrue {
		fmt.Println("1")
	} else if myNum == 101 || isTrue {
		fmt.Println("2")
	} else if myNum > 1000 && isTrue == false {
		fmt.Println("3")
	}
}
