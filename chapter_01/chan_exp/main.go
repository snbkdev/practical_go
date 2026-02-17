// Использование канала
package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <- c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 4, 5, 6, 2, 1, 7, 0, 9}
	go printCount(c)
	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond * 10)
	fmt.Println("End of main")	 // 8 4 5 6 2 1 7 0 9 End of main
}