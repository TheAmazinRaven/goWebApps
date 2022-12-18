package main

import "fmt"

func main() {
	firstLine := "I brush against those freckles that I hated so"

	for i, l := range firstLine {
		fmt.Println(i, l)
	}
}
