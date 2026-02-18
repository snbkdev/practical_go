// CLI-приложение hello world
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCommand *cobra.Command

func init() {
	helloCommand = &cobra.Command{
		Use: "hello",
		Short: "print hello world",
		Run: sayHello,
	}

	helloCommand.Flags().StringP("name", "n", "world", "who to say hello to")
	helloCommand.MarkFlagRequired("name")
	helloCommand.Flags().StringP("language", "l", "en", "which language to say hello in")
}

func sayHello(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	greeting := "hello"
	language, _ := cmd.Flags().GetString("language")

	switch language {
	case "en":
		greeting = "hello"
	case "sp":
		greeting = "hola"
	case "fr":
		greeting = "bonjour"
	case "de":
		greeting = "hallo"
	}

	fmt.Printf("%s %s! \n", greeting, name)
}

func main() {
	helloCommand.Execute()
}