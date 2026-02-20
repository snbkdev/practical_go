// Использование каналов
package main

import (
	"os"
	"time"
)

func main() {
	echo := make(chan []byte)

	go readStdin(echo)

	for {
		select {
		case buf := <- echo:
			os.Stdout.Write(buf)
		case <- time.After(30 * time.Second):
			break
		}
	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}