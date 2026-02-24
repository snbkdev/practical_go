// Отправка данных формы
package main

import (
	"fmt"
	"net/http"
	"time"
)

type comment struct {
	username string
	text string
	dateString string
}

var comments []comment

func commentHandler(w http.ResponseWriter, r *http.Request) {
	body := `<html><head></head><body>`
	commentBody := ""
	for i := range comments {
		commentBody += fmt.Sprintf("<div class='comment'>%s (%s) - @%s</div>", comments[i].text, comments[i].dateString, comments[i].username)
	}

	body += fmt.Sprintf(`<h1>Comments</h1>
	%s
	<form method="POST" action="/comments"><div><input type="text" placeholder="Username" name="username" /></div>
	<textarea placeholder="Comment text" name="comment"></textarea>
	<div><input type="submit" value="Submit" /></div></form></body></html>`, commentBody)

	w.Write([]byte(body))
}

func main() {
	http.HandleFunc("GET /comments", commentHandler)
	http.HandleFunc("POST /comments", postHandler)

	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}

// Установка cookie
func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")

	usernameCookie, err := r.Cookie("username")
	if err == nil {
		username = usernameCookie.Value
	}

	commentText := r.Form.Get("comment")

	comments = append(comments, comment{
		username: username,
		text: commentText,
		dateString: time.Now().Format(time.RFC3339),
	})

	http.SetCookie(w, &http.Cookie{
		Name: "username",
		Value: username,
		Expires: time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "/comments", http.StatusNotFound)
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	username := ""
	usernameCookie, err := r.Cookie("username")
	if err == nil {
	username = usernameCookie.Value
	}
	body := `
	<html>
		<head>
			<title>Comments</title>
			<style type="text/css">
				body {
					width: 500px;
					margin: 0 auto;
				}
				h1 {
					margin: 0;
					padding: 0;
				}
				div {
					padding: 20px 0;
				}
				textarea, input[type="text"] {
					width: 100%;
				}
				textarea {
					height: 200px;
				}
				.comment {
					padding: 10px;
					border: 1px solid #ddd;
					margin-bottom: 4px;
				}
			</style>
		</head>
	<body>`

	commentBody := ""
	for i := range comments {
		displayName := comments[i].username
		if username != "" && displayName == username {
			displayName = "You"
		}
		commentBody += fmt.Sprintf("<div class='comment'>%s (%s) - @%s</div>", comments[i].text, comments[i].dateString, displayName)
	}
	body += commentBody
	body += `</body></html>`

	w.Write([]byte(body))
}