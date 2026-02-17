package main

import "fmt"

func getString() (string, string) {
	return "foo", "bar"
}

func main() {
	n1, n2 := getString()
	fmt.Println(n1, n2)

	n3, _ := getString()
	fmt.Println(n3)
}