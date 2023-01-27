package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("hello from go client")
	GetRequest()
}

func GetRequest() {
	const url = "http://localhost:8000/get"

	// TODO => fetch a GET request
	res, err := http.Get(url)

	// TODO => handle the error
	HandleNilErrors(err)

	// TODO => close the request to the server
	defer res.Body.Close()

	// TODO => manage the received response data
	fmt.Println("The Status Code -> ", res.StatusCode)
	fmt.Println("The Content Length -> ", res.ContentLength)

	// TODO => read the response body and parse it
	ContentByteData, err := ioutil.ReadAll(res.Body)
	HandleNilErrors(err)
	// convert it from byte data to normal readable string
	fmt.Println("The Response Body -> ", string(ContentByteData))
}

func HandleNilErrors(err error) {
	if err != nil {
		panic(err)
	}
}
