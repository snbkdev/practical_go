// Определение допустимых параметров с помощью перечисления
package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "world", "a name to say hello to")

type language = string

var userLanguage language

const (
	English = "en"
	Spanish = "sp"
	French = "fr"
	German = "de"
)

func init() {
	flag.StringVar(&userLanguage, "lang", "en", "language to use (en, sp, fr, de)")
	flag.Parse()
}

func main() {
	switch(userLanguage) {
	case English:
		log.Printf("Hello %s! \n", *name)
	case Spanish:
		log.Printf("Hola %s! \n", *name)
	case French:
		log.Printf("Bonjour %s! \n", *name)
	case German:
		log.Printf("Hallo %s! \n", *name)
	}
}