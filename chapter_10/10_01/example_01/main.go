// Раздача файлов с помощью пакета http
package main

import "net/http"

func main() {
	dir := http.Dir("./files")
	if err := http.ListenAndServe(":8080", http.FileServer(dir)); err != nil {
		panic(err)
	}
}