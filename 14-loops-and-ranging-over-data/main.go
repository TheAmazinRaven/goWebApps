package main

import "fmt"

func main() {
	animals := []string{"dog", "fish", "cat", "bear", "cow", "horse"}

	for i, animal := range animals {
		fmt.Println(i, animal)
	}
}
