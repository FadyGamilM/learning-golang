package main

import "fmt"

func main() {
	res := adder(5, 6)
	fmt.Println(res)

	res = proAdder(5, 6, 1, 8)
	fmt.Println(res)
}

func adder(a int, b int) int {
	return a + b
}

// TODO => func that has no clew about how many var will be sent
func proAdder(nums ...int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
