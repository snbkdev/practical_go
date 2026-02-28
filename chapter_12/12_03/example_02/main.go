// декодирование данных JSON с помощью инструмента Sonic
package main

import (
	"log"
	"strings"

	"github.com/bytedance/sonic"
)

func main() {
	var receiver []map[string]interface{}

	jsonData := strings.NewReader(`[
		{"email": "inigo@example.com", "name": "Inigo"},
		{"email": "armor_king@tekken.asia", "Armor King"}
	]`)

	decoder := sonic.ConfigDefault.NewDecoder(jsonData)
	decoder.Decode(&receiver)

	for k := range receiver{
		log.Println(receiver[k])
	}
}