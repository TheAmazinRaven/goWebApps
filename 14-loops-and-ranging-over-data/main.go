package main

import "fmt"

func main() {
	animals := make(map[string]string)

	animals["dog"] = "Sheba"
	animals["dog2"] = "Prince"

	for animalType, animal := range animals {
		fmt.Println(animalType, animal)
	}
}
