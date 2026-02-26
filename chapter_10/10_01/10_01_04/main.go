// Передача URL-адреса в шаблон
package main

import (
	"flag"
	"html/template"
	"net/http"
)

var t *template.Template
var l = flag.String("location", "http://localhost:8080", "A location")

var tpl = `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>A demo</title>
				<link rel="stylesheet" href="{{.Location}}/styles.css">
			</head>
			<body>
				<p>A Demo !!!</p>
			</body>
			</html>`

func servePage(w http.ResponseWriter, r *http.Request) {
	data := struct{ Location *string }{Location: l,}
	t.Execute(w, data)
}