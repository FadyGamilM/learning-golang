package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Id       int      `json:"id"`
	Name     string   `json:"courseName"`
	Price    string   `json:"price"`
	HashCode string   `json:"-"`
	Tags     []string `json:"tags"`
}

func main() {
	fmt.Println("Lets create json data .. ")
	// EncodeDataToJson()
	ConsumeJsonData()
}

func EncodeDataToJson() {
	// TODO => create some seed data
	data := []Course{
		{Id: 1, Name: "Master React.js", Price: "$99.9", HashCode: "34g53asg3as21g", Tags: []string{"frontend", "web development"}},
		{Id: 2, Name: "golang for dummies", Price: "$205.9", HashCode: "454asfa3s1323fs", Tags: []string{"backend", "web development"}},
		{Id: 3, Name: "20 MERN stack project", Price: "$999.9", HashCode: "8956a5sfxvehvn", Tags: []string{"full stack", "web development"}},
	}

	// TODO => convert this data into json data
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", jsonData)
}

func ConsumeJsonData() {
	// prepare some fake received json data
	jsonData := []byte(`
      {
         "id": 3,
         "courseName": "20 MERN stack project",
         "price": "$999.9",
         "tags": ["full stack", "web development"]
      }
   `) // the data is of type slice of bytes because json.Valid deals with bytes slice and also when we read the data we read it as []byte remember !

	// TODO => we need to ensure that this is a valid json data
	//!SECTION => you might create struct for each res
	var parsedData Course

	//!SECTION => you might also want to let the map handle this for you
	var parsedData_v2 map[string]interface{} // we are sure that the key is coming as string but the value might be slice "array" so we use interface

	isValid := json.Valid(jsonData)
	if !isValid {
		fmt.Println("JSON data is not valid")
		return
	}
	// i gave it a ref to my struct instance because when we pass a param to a method we know that its passed by value but we need to store the data into that instance so i sent the ref
	json.Unmarshal(jsonData, &parsedData)
	fmt.Printf("Version 1 -> \n %#v \n", parsedData)

	json.Unmarshal(jsonData, &parsedData_v2)
	fmt.Printf("Version 2 -> \n %#v \n", parsedData_v2)
	for key, val := range parsedData_v2 {
		fmt.Printf("Key => %v, Value => %v, Type => %T  \n", key, val, val)
	}
}

// you might don't want to create a struct for each json response you will receive and instead store it into a map
func ConsumeJsonData_V2() {
}
