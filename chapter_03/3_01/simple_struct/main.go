package main

import "fmt"

type configuration struct {
	Enabled bool
	Path string
}

func main() {
	animal := struct {
		name string
		speak func() string
	} {
		name: "cat",
		speak: func() string {
			return "meow"
		},
	}

	fmt.Println(fmt.Sprintf("our animal's name is %s and it says %s", animal.name, animal.speak()))
}