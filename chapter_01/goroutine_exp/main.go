package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 5)
	}
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 20)
	fmt.Println("Hello world")
	time.Sleep(time.Millisecond * 10)
}