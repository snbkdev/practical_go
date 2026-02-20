package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("checking for environment variable CLEANUP")

	if envvar := os.Getenv("CLEANUP"); envvar != "" {
		fmt.Println("did not find it, value is: ", envvar)
	}
}