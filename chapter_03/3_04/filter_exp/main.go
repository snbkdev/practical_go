// Реализация функции filter
package main

import (
	"fmt"
	"unicode"
)

func filter[T any](items []T, fx func(T) bool) []T {
	var filtered []T
	for _, v := range items {
		if fx(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func main() {
	strings := []string{"my", "name", "is", "marshall", "law"}

	strings = filter[string](strings, func(s string) bool {
		return unicode.IsUpper(rune(s[0]))
	})

	fmt.Println(strings)
}