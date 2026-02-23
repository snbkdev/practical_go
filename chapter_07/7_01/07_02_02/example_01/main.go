// Создание файла и запись данных в него
package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("test")
}