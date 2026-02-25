// Передача срезов/массивов в шаблон
package main

import (
	"log"
	"net/http"
	"html/template"
)

type comment struct {
	Username string
	Text string
}

type Page struct {
	Title, Content string
	Comments []comment
}

var t = template.New("templates")

func routeComments(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An Example",
		Content: "Have fun stormin da castle",
		Comments: []comment{
			{Username: "Bill", Text: "Looks like a good example"},
			{Username: "Jill", Text: "I really enjoyed this article"},
			{Username: "Phil", Text: "I do not like to read"},
		},
	}

	if err := t.ExecuteTemplate(w, "index.html", p);
	err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func init() {
	_, err := t.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Error loading templates:" + err.Error())
	}
}

func main() {
	http.HandleFunc("/comments", routeComments)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}