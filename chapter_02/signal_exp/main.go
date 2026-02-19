// Корректное завершение работы с использованием сигналов
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	handleFunc := newHandler()
	server := &http.Server{
		Addr: ":8080",
		Handler: handleFunc,
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)

	go func() {
		server.ListenAndServe()
	}()

	time.Sleep(5 * time.Second)
	<- ch
	if err := server.Shutdown(nil); err != nil {
		panic(err)
	}
}

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Armor King"
	}

	fmt.Fprint(w, "Hello, my name is ", name)
}