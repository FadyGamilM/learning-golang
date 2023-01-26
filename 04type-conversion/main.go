package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your rate btn 1 and 5")

	input, _ := reader.ReadString('\n')

	fmt.Printf(" %T ", input)
	fmt.Println(input)

	// now lets say you need to do some calc on the user input
	// the used package is `strconv`
	// we use strings.TrimSpace because the input will contains both he number you entered attached to the \n due to your enter press
	result, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// check if the error contains something ..
	if err != nil {
		fmt.Println("Error happens ", err)
	} else {
		result := result + 1
		fmt.Println("your rating is -> ", result)
	}
}
