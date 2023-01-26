package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package")

	currTime := time.Now()
	fmt.Println(currTime) // 2023-01-26 13:39:45.5065542 +0200 EET m=+0.012491001

	// formatting the time by giving it a reference to the format you need
	fmt.Println(currTime.Format("01-02-2006"))                 // exactly this specific date
	fmt.Println(currTime.Format("01-02-2006 Monday"))          // exactly this specifc day
	fmt.Println(currTime.Format("01-02-2006 15:04:05 Monday")) // exactly this specific hour

	// to create a new date var
	myDate := time.Date(2023, time.December, 1, 10, 0, 0, 0, time.UTC)
	fmt.Println(myDate.Format("01-02-2006"))
}
