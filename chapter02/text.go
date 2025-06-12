package main

import "fmt"

func main() {
	asString := "Hello World! €"
	fmt.Println("First character: ", string(asString[0]))

	r := '€'
	fmt.Println("As an int32 value:", r)

	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	for _, v := range asString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	for _, v := range asString {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}
