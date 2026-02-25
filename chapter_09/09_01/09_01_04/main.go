// Буферизация результата выполнения шаблона
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./templates/index.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An example",
		Content: "greate example",
	}

	var b bytes.Buffer	 // Создает буфер для хранения результата выполнения шаблона
	err := t.Execute(&b, p)
	if err != nil {
		fmt.Fprint(w, "A error ocured") 	// Обрабатывает ошибки, возникшие при выполнении шаблона
		return
	}

	b.WriteTo(w) 	// Копирует сождержимое буфера в поток вывода
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}