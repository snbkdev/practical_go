// Утечка контекста
package main

import (
	"context"
	"fmt"
	"net/http"
)

type config struct {
	HomepageDescription string
	Pageviews           int64
}

func main() {
	c := config{
		HomepageDescription: "my 1997-style personal web site",
		Pageviews:           0,
	}

	ctx, _ := context.WithCancel(context.Background()) // создает контекст
	ctx = context.WithValue(ctx, "webConfig", c)       // добавляет конфигурацию в контекст в качестве значения

	http.HandleFunc("/home", homeHandler(ctx)) // извлекает значение из контекста и приводит его к типу конфигурации
	http.HandleFunc("/guestbook", guestbookHandler(ctx))
	http.ListenAndServe(":8081", nil)
}

func homeHandler(ctx context.Context) http.HandlerFunc {
	myValue := ctx.Value("webConfig").(config)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fmt.Sprintf("welcome to %s", myValue.HomepageDescription))
	}
}

func guestbookHandler(ctx context.Context) http.HandlerFunc {
	myValue := ctx.Value("webConfig").(config)
	return func(w http.ResponseWriter, r *http.Request) {
		myValue.Pageviews++
		fmt.Fprintln(w, fmt.Sprintf("welcome to my guestbook. hit counter since server start: %v", myValue.Pageviews))
	}
}

// go vet main.go
// Результат - chapter_06/6_02/main.go:21:7: the cancel function returned by context.WithCancel should be called, not discarded, to avoid a context leak
