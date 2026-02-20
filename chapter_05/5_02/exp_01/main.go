// Использование горутины для выполнения задачи
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("type anything below for up to 30 seconds")

	go echo(os.Stdin, os.Stdout)

	time.Sleep(30 * time.Second)

	fmt.Println("Timed out")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}