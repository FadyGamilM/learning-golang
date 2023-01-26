package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the rating section")

	// define a reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your rating for our service")

	// comma Ok syntax [err ok syntax]
	// no try.catch here in go
	input, _ := reader.ReadString('\n') // keep read until user hits enter

	fmt.Println("Thanks for your rating ", input)
	fmt.Printf("%T", input) // its a string !!
}
