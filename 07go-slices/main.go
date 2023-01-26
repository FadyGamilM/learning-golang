package main

import (
	"fmt"
	"sort"
)

func main() {
	// TODO => define slices way [1]
	usernames := []string{}
	usernames = append(usernames, "magy")
	fmt.Println(usernames)

	// TODO => define slices way [2]
	ids := make([]int, 1) // initialize a memory space with size of 1 int item
	// -> Now a new memory will be reallocated ..
	ids = append(ids, 5)
	ids = append(ids, 2)
	ids = append(ids, 3)
	fmt.Println(ids)

	// TODO => to sort the items
	fmt.Println(sort.IntsAreSorted(ids)) // to know if they are sorted or not ..
	sort.Ints(ids)
	fmt.Println(ids)
	fmt.Println(sort.IntsAreSorted(ids)) // to know if they are sorted or not ..

	// TODO => how to slice a slice ..!
	names := []string{}
	names_v1 := []string{}
	names_v2 := []string{}
	names = append(names, "fady", "magy", "roshdy", "samy", "injy")
	fmt.Println(names)

	names_v1 = append(names[:3]) // get the 0th, 1st, and 2nd and don't take the 3rd
	fmt.Println(names_v1)

	names_v2 = append(names[1:3]) // ignore the 0th and go for 1st and 2nd and stop and don't take the 3rd
	fmt.Println(names_v2)

	// TODO => how to remove item by its index
	items := []int{}
	items = append(items, 0, 1, 2, 3, 4, 5, 6)
	idx := 2 // index we need to remove the item at it
	items = append(items[:idx], items[idx+1:]...)
	fmt.Println(items)
}
