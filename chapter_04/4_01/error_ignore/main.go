// Игнорирование возвращаемых ошибок
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}

	return strings.Join(parts, " "), nil
}

func main() {
	args := os.Args[1:]
	result, _ := Concat(args...)
	fmt.Printf("Concatenated strings - %s\n", result)
}