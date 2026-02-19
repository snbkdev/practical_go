// Конфигурация на основе переменных окружения
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port string
	if port =os.Getenv("PORT"); port == "" {
		panic("environment variable PORT has not been set!")
	}

	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":"+port, nil)
}

func HomePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	fmt.Fprint(res, "The HomePage")
}