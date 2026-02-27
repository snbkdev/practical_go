package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{
	"firstName": "Jean", "lastName": "Bartik", "age": 86, "education": [
	{"institution": "Northwest Missouri State Teachers College", "degree": "Bachelor of Science in mathematics"}, 
	{"institution": "University of Pennsylvania", "degree": "Masters in English"}],
	"spouse": "William Bartik", "children":["Timothy John Bartik", "Jane Helen Bartik", "Mary Ruth Bartik"]}`)

func main() {
	var f interface{}

	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})

	fmt.Println(m["firstName"])
}