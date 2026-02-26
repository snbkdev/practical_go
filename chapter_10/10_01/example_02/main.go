// Раздача файла с помощью пользовательского обработчика
package main

import "net/http"

func main() {
	http.HandleFunc("/", readme)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func readme(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./files/readme.txt")
}