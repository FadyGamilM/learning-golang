package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("hello from go client")
	GetRequest()
	PostRequest()
	PostFormRequest()
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
	fmt.Println("The Response of the GET request is -> ", string(ContentByteData))
}

func PostRequest() {
	const url = "http://localhost:8000/post"

	// TODO => fake json payload for my request body
	reqBody := strings.NewReader(`
      {
         "username" : "magy",
         "age" : 25
      }
   `)

	// TODO => send the request
	res, err := http.Post(url, "application/json", reqBody)
	HandleNilErrors(err)

	// TODO => close the request connection
	defer res.Body.Close()

	// TODO => read the response
	ContentDataByte, err := ioutil.ReadAll(res.Body)
	HandleNilErrors(err)
	fmt.Println("The Response of the POST request is -> ", string(ContentDataByte))
}

func PostFormRequest() {
	// HINT => this type of requests is receiving a url-encoded data

	const server_url = "http://localhost:8000/postform"

	// TODO => get the data ready
	data := url.Values{}
	data.Add("id", "1")
	data.Add("firstname", "fady")
	data.Add("lastname", "gamil")

	// TODO => send the requst
	res, err := http.PostForm(server_url, data)
	HandleNilErrors(err)

	// TODO => close the connection
	defer res.Body.Close()

	// TODO => handle the response
	contentBytes, err := ioutil.ReadAll(res.Body)
	HandleNilErrors(err)
	fmt.Println("The Response of the POST request is -> ", string(contentBytes))
}

func HandleNilErrors(err error) {
	if err != nil {
		panic(err)
	}
}
