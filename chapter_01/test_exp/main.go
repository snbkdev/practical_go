package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reverseName(name string) string {
	reversed := make([]byte, 0)
	for i := len(name) - 1; i >= 0; i-- {
		reversed = append(reversed, name[i])
	}

	return string(reversed)
}

func main() {
	fmt.Print("Enter Your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("could not read from stdin", err)
	}
	name = strings.TrimSpace(name)

	reversed := reverseName(name)
	fmt.Println(reversed, ", olleH")
}

func reverseNameFixed(name string) string {
	reversed := make([]rune, 0)
	runes := []rune(name)
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}

	return string(reversed)
}