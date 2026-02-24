// Ручная маршрутизация на основе метода с дополнительной логикой
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type comment struct {
	ID int `json:"id,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func upsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var postComment comment
		if err := json.Unmarshal(postBody, &postComment); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if postComment.ID == 0 {
			// createCommentHandler(w, r)
		} else {
			// upsertCommentHandler(w, r)
		}
	}

	if r.Method == http.MethodPut || r.Method == http.MethodPatch {
		// upsertCommentHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/comments/upsert", upsertHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("could not start server")
	}
}