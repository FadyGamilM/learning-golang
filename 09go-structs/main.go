package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Name string
	Age  float32
	Role string
}

func main() {
	fady := User{"fady", 25, "admin"}
	// TODO => to print the instance in details use +v not v
	fmt.Printf("the details are %+v \n", fady)
	// TODO => you can access specific field
	fmt.Printf("name is %v \n", fady.Name)

	// TODO => working with randmoness in golang
	rand.Seed(time.Now().UnixNano()) // unixNano is an int64 number
	diceNumber := rand.Intn(6) + 1   // so if we got 0 its 1 and if we got 5 its a 6 ..
	fmt.Println(diceNumber)
}
