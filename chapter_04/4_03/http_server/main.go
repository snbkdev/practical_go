package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(fmt.Sprintf("could not start server %v", err))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("fake panic!"))
}