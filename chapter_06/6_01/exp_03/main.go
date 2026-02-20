package main

import (
	"encoding/json"

	"fmt"
)

type user struct {
	username string `json:"username"`
	Email string `json:"email"`
	firstName string
}

func main() {
	m := user{
		username: "Test",
		Email: "example@example.com",
		firstName: "Testov",
	}

	out, err := json.Marshal(m) // Функция маршалинга выполняется без ошибок
	if err != nil {
		panic("could not marshal")
	}

	fmt.Println(string(out))
}

// go vet main.go
// Результат: struct field username has json tag but is not exported