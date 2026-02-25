// Использование простого HTML-шаблона
package main

import (
	"net/http"
	"html/template"
)

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An Example",
		Content: "Have fun stormin da castle",
	}

	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}