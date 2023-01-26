package main

import "fmt"

func main() {
	fady := User{1, "fady", "fady@gmail.com", true}
	isActive := fady.ValidateEmail()
	fmt.Println(isActive)

	// TODO => user is passed by value [a copy] not by ref
	fmt.Printf("Before de-activating the status : %v \n", fady.Active)
	fady.DeActivateUser()
	fmt.Printf("After de-activating the status : %v \n", fady.Active)
}

type User struct {
	Id       int
	Username string
	Email    string
	Active   bool
}

// methods on this User strcut
func (user User) ValidateEmail() bool {
	// TODO : some work for email valdiation and now return that its a valid email
	return true
}

func (user User) DeActivateUser() {
	// this method make his/her status to be false
	user.Active = false
	fmt.Printf("inside deactivation method : %v \n", user.Active)
}
