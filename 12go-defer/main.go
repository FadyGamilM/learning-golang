package main

import "fmt"

func main() {
	// the program will do the following
	/*
					   => print hello
					   => print world
					   => open stack
					      -> push(golang) : top -> golang -> nil
					      -> push(is) : top -> is -> golang -> nil
					      -> push(amazing) : top -> amazing -> is -> golang -> nil
				      => print Bye
		            => execute the function [its not defered]
		               -> push(0) : top -> 0 -> amazing -> is -> golang -> nil
		               -> push(1) : top -> 1-> 0 -> amazing -> is -> golang -> nil
		               -> push(2) : top -> 2-> 1-> 0 -> amazing -> is -> golang -> nil
		               -> push(3) : top -> 3->  2-> 1->  0 -> amazing -> is -> golang -> nil
		               -> push(4) : top -> 4 -> 3->  2-> 1-> 0 -> amazing -> is -> golang -> nil
				      => now empty the stack from the top
				         -> print 4 -> 3 -> 2 -> 1 -> 0 -> amazing -> then print is -> then print golang
	*/
	fmt.Println("hello")
	fmt.Println("world")
	defer fmt.Println("golang")
	defer fmt.Println("is")
	defer fmt.Println("amazing")
	fmt.Println("Bye")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
