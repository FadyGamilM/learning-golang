package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	filePath := "./myfile.txt"
	content := "learning golang is an amazing skill added to my toolbox"
	CreateAndWriteToFile(filePath, content)
	ReadFromFile(filePath)
}

func CreateAndWriteToFile(filePath string, content string) {
	// TODO => Create the file
	file, err := os.Create(filePath)
	HandleNilError(err)

	// TODO => write to the file
	length, err := io.WriteString(file, content)
	HandleNilError(err)
	fmt.Printf("The length of the written content is -> %v \n", length)

	// TODO => finally close the file
	defer file.Close()
}

func ReadFromFile(filePath string) {
	dataBytes, err := ioutil.ReadFile(filePath)
	HandleNilError(err)
	fmt.Printf("The data from the file -> \n")
	fmt.Printf("%v \n", string(dataBytes))
}

func HandleNilError(err error) {
	if err != nil {
		panic(err)
	}
}
