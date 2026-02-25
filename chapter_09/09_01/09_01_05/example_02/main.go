// Использование механизма наследования шаблонов
package main

import (
	"html/template"
	"net/http"
)

var t map[string]*template.Template

func init() {
	t = make(map[string]*template.Template)

	temp := template.Must(template.ParseFiles("templates/index.html", "templates/title.html"))
	t["title.html"] = temp

	temp = template.Must(template.ParseFiles("templates/index.html", "templates/content.html"))
	t["content.html"] = temp
}

type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An example 02",
		Content: "Text for second example",
	}
	t["content.html"].ExecuteTemplate(w, "base", p)
}

func dispayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "armor_king",
		Name: "Armor King",
	}

	t["title.html"].ExecuteTemplate(w, "base", u)
}

func main() {
	http.HandleFunc("/user", dispayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}