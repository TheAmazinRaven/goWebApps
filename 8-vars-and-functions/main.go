package main

import "fmt"

func main() {
	fmt.Println("Hello, sis!")

	var whatToSay string
	whatToSay = "girrrrl"
	var i int
	i = 29

	fmt.Println(whatToSay)
	fmt.Println("I is set to", i)

	whatWasSaid, theOtherThingSaid := saySomething()

	fmt.Println(whatWasSaid, theOtherThingSaid)

}

func saySomething() (string, string) {
	return "I am a", "magickal gyal"
}
