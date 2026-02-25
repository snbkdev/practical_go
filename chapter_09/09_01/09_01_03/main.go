// Кэширование парсированного шаблона
package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("templates/index.html"))

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: "An Example",
		Content: "Have fun stormin de castle",
	}

	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}