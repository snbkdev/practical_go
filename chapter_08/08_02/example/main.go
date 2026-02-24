package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type comment struct {
	text       string
	dateString string
}

var comments []comment

func getComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if commentID == 0 || len(comments) < commentID {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Command %d : %s", commentID, comments[commentID-1].text)
}

func main() {
	http.HandleFunc("GET /comments", getComments)
	http.HandleFunc("GET /comments/{id}", getComment)
	http.HandleFunc("POST /comments", postComments)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

func getComments(w http.ResponseWriter, r *http.Request) {
	commentBody := ""
	for i := range comments {
		commentBody += fmt.Sprintf("%s (%s)\n", comments[i].text, comments[i].dateString)
	}

	fmt.Fprintln(w, fmt.Sprintf("Comments: \n%s", commentBody))
}

func postComments(w http.ResponseWriter, r *http.Request) {
	commentText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comments = append(comments, comment{
		text:       string(commentText),
		dateString: time.Now().Format(time.RFC3339),
	})

	w.WriteHeader(http.StatusOK)
}
