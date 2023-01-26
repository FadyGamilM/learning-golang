package main

import "fmt"

func main() {
	// TODO => to define a map using make
	student_grade := make(map[string]float32)

	student_grade["fady"] = 3.7
	student_grade["ramy"] = 3.5

	// TODO => to loop through map items
	for key, val := range student_grade {
		fmt.Printf("Student %v,  got grade -> %v \n", key, val)
	}

	// TODO => we can delete an item from map using delete method
   fmt.Println("After removing `fady`")
	delete(student_grade, "fady")
	for key, val := range student_grade {
		fmt.Printf("Student %v,  got grade -> %v \n", key, val)
	}
}
