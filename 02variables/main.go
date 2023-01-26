package main

import "fmt"

// you can't define a variables outside the func with way 3 ..
//!ERROR: token := "gasgasg251345"

// capital first char is used to make this public var
const LoginSecretKey string = "ahashsah4534351cac3"

func main() {
	fmt.Println("+ Way 1 of defining a variable")
	// [var] [varName] [datatype]
	var username string = "fady"
	fmt.Println(username)

	fmt.Println("+ Way 2 of defining a variable")
	username2 := "ahmed"
	fmt.Println(username2)

	fmt.Println("+ [recommended] Way 3 of defining a variable")
	username3 := "ahmed"
	fmt.Println(username3)

	// if we need to know the type of the variable
	fmt.Printf("+ variable is of type: %T \n", username)

	// default values, golang doesn't put rubbish in the uninitialized vars
	var age byte
	fmt.Println("+ the default value of age is -> ", age)
	fmt.Printf("%T \n", age)

	fmt.Println("+ you can access the public cosnstant vars here")
	fmt.Println(LoginSecretKey)
}
