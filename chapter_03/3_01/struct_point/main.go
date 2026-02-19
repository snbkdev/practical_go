// Метод с указателем в качестве получателя
package main

import "fmt"

type character struct {
	name string
}

func (ch *character) fixName() {
	ch.name = "Paul Phoenix"
}

func main() {
	ch := new(character)
	ch.name = "Marshal Low"
	fmt.Println("my name is ", ch.name)

	ch.fixName()
	fmt.Println("just kidding, my name is ", ch.name)
}