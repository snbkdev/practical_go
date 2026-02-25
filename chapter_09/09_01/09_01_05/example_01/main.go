// Использование вложенных шаблонов
package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("templates/index.html", "templates/head.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Example 01",
		Content: "Some text...",
	}

	t.ExecuteTemplate(w, "templates/index.html", p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
