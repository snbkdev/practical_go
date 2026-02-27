package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func displayPage(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("accept")
	var err error
	var b []byte
	var ct string

	switch t {
	case "application/vnd.mytodos.json; version=2.0":
		data := testMessageV2{"Version 2"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=2.0"

	case "application/vnd.mytodos.json; version=1.0":
		fallthrough
	default:
		data := testMessageV1{"Version 1"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=1.0"
	}

	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	w.Header().Set("Content-Type", ct)
	fmt.Fprint(w, string(b))
}

type testMessageV1 struct {
	Message string `json:"message"`
}

type testMessageV2 struct {
	Info string `json:"info"`
}

func main() {
	ct := "application/vnd.mytodos.json; version=2.0"

	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	req.Header.Set("Accept", ct)

	res, _ := http.DefaultClient.Do(req)
	if res.Header.Get("Content-Type") != ct {
		fmt.Println("unexpected content type returned")
		return
	}

	b, _ := io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", b)

	// http.HandleFunc("/test", displayPage)
	// http.ListenAndServe(":8080", nil)
}