// Извлечение дополнительной информации о файле
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("structured.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("file: name is %s, mode is %v, size is %d. Is directory %v", info.Name(), info.Mode(), info.Size(), info.IsDir()))

	lineJSON := make(map[string]interface{})
	var bChunk []byte

	for {
		b := make([]byte, 2)
		_, err := file.Read(b)
		if err != nil {
			break
		}

		bChunk = append(bChunk, b[0:]...)

		if err := json.Unmarshal(bChunk, &lineJSON); err == nil {
			log.Println(lineJSON)
			bChunk = []byte{}
		}
	}
}