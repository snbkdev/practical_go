// Добавление функции в шаблон
package main

import (
	"html/template"
	"net/http"
	"time"
)

var tmpl = `<!DOCTYPE HTML><html><head><meta charset="utf-8">
<title>Data Example</title></head><body><p>{{.Date | dateformat "Jan 2, 2006"}}</p></body></html>`
var funcMap = template.FuncMap{
	"dateformat": dateFormat,
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.New("date")
	t.Funcs(funcMap)
	t.Parse(tmpl)

	data := struct{ Date time.Time } {
		Date: time.Now(),
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}