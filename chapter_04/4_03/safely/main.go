package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/cookoo/safely"
)

func message() {
	fmt.Println("inside goroutine")
	panic(errors.New("Oops!!!"))
}

func main() {
	safely.Go(message)
	fmt.Println("outside goroutine")
	time.Sleep(1000)
}