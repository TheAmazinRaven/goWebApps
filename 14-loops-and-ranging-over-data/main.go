package main

import "fmt"

func main() {
	animals := []string{"dog", "fish", "cat", "bear", "cow", "horse"}

	for _, animal := range animals {
		fmt.Println(animal)
	}
}
