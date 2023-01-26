package main

import "fmt"

func main() {
	myAge := 25

	agePtr := &myAge

	fmt.Println("Reference to my age is -> ", agePtr)
	fmt.Println("Actual value to my age is -> ", *agePtr)
}
