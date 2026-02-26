// Простой пользовательский HTTP-клиент
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	cc := &http.Client{Timeout: time.Second}
	res, err := cc.Get("http://www.manning.com")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Printf("%s", b)
}