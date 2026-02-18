// Приложение Hello World, использующее флаги командной строки
package main

import (
	"flag"
	"fmt"
)


var name = flag.String("name", "world", "A name to say hello to")
var inSpanish bool

func init() {
	flag.BoolVar(&inSpanish, "spanish", false, "use spanish language")
	flag.BoolVar(&inSpanish, "s", false, "use spanish language")
	flag.Parse()
}

func main() {
	if inSpanish {
		fmt.Printf("Hola %s\n", *name)
	} else {
		fmt.Printf("Hello %s !\n", *name)
	}
}