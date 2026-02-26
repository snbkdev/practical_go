// Раздача подкаталога
package main

import "net/http"

func main() {
	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))

	http.Handle("/static", handler)
	http.HandleFunc("/", homePage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	//
}