// Объединение шаблонов
package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("templates/index.html", "templates/quote.html"))
}

type Page struct {
	Title string
	Content template.HTML
}

type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{
		Quote: `You keep using thatword. I do not think it means what you think it means`,
		Person: "Armor King",
	}

	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String())

	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "A user",
		Content: qc,
	}

	t.ExecuteTemplate(w, "index.html", p)
}