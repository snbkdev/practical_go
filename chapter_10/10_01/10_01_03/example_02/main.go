package main

import (
	"fmt"
	"log"
)

var myString string

func main() {
	log.Println(fmt.Sprintf("embedded value: %s", myString))
}