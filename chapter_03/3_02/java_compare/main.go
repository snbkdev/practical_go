// Go
package main

import "fmt"

type Animal interface {
	speak()
}

type Cat struct {}

func (c Cat) speak() {
	fmt.Println("meow")
}

func NewCat() *Cat {
	return &Cat{}
}

type Dog struct {}

func (d Dog) speak() {
	fmt.Println("woof")
}

func NewDog() *Dog {
	return &Dog{}
}

type Llama struct {}

func NewLLama() *Llama {
	return &Llama{}
}

func main() {
	var a Animal

	c := NewCat()
	a = c
	a.speak()

	d := NewDog()
	a = d
	a.speak()

	// l := NewLLama()
	// a = l
	// result -> cannot use l (type *Llama) as type Animal in assignment: *Llama does not implement Animal (missing speak method)

	l := NewLLama()
	a = l
	a.speak()
}

func (l Llama) speak() {
	fmt.Println("nondescript animal noise")
}