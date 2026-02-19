// Возврат ошибки
package main

import (
	"errors"
	"strings"
)

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}

	return strings.Join(parts, " "), nil
}