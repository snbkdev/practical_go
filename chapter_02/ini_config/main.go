// Парсинг конфигурационного файла в формате INI
package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	config, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(config.Section("Section").Key("path").String())
	enabled, err := config.Section("Section").Key("enabled").Bool()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(enabled)
}