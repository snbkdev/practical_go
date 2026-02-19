// Анонимные идентификаторы
package main

import (
	"fmt"
)

type Animal struct {
	string
}

func (a Animal) speak() string {
	return a.string
}

func main() {
	a := Animal{
		"cat",
	}

	fmt.Println(a.speak())

	a.string = "dog"
	fmt.Println(a.speak())
}
