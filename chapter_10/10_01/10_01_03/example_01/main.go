// Встраивание файлов в исполняемый файл с помощью go:embed
package main

import (
	"embed"
	"net/http"
)

var f embed.FS

func main() {
	if err := http.ListenAndServe(":8088", http.FileServer(http.FS(f))); err != nil {
		panic(err)
	}
}