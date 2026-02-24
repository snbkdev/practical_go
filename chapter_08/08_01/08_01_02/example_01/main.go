// Управление сервером
package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("you should never see me"))
}

func main() {
	muxer := http.NewServeMux()
	muxer.HandleFunc("GET /timeout", timeoutHandler)

	server := http.Server{
		Addr: ":8080",
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler: muxer,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("could not start server: %s", err.Error()))
	}
}