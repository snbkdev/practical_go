// Передача состояния через структуру
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type serverControl struct {
	db *pgx.Conn
}

func main() {
	var sc serverControl
	{
		var err error
		sc.db, err = pgx.Connect(context.Background(), "postgres//localhsot:1234")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to databaseL %v\n", err)
			os.Exit(1)
		}
	}

	http.HandleFunc("GET /database", sc.databaseHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type comment struct {
	UserID int `json:"userID"`
	Comment string `json:"comment"`
}

func (sc serverControl) databaseHandler(w http.ResponseWriter, r *http.Request) {
    var comments []comment
    
    rows, err := sc.db.Query(context.Background(), `select user_id, comment from comments limit $1`, 5)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer rows.Close()
    
    for rows.Next() { 
        var c comment
        if err := rows.Scan(&c.UserID, &c.Comment); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        comments = append(comments, c)
    }
    
    if err = rows.Err(); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
    output, err := json.Marshal(comments)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(output)
}