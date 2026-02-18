// Простой калькулятор, реализованный в виде CLI-приложения

package main

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func operation(op, n1, n2 string) (float32, error) {
	num1, err := strconv.Atoi(n1)
	if err != nil {
		return 0, err
	}

	num2, err := strconv.Atoi(n2)
	if err != nil {
		return 0, err
	}

	fl1 := float32(num1)
	fl2 := float32(num2)
	switch op {
	case "add":
		return fl1 + fl2, nil
	case "sub":
		return fl1 - fl2, nil
	case "mul":
		return fl1 * fl2, nil
	case "div":
		return fl1 / fl2, nil
	}

	return 0, nil
}

var cmdAdd = &cobra.Command{
	Use: "add",
	Short: "add two numbers",
	Long: "add two numbers add <number1> <number2>",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := operation("add", args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("result: %f\n", result)
	},
	Args: cobra.ExactArgs(2),
}

var cmdSub = &cobra.Command{
	Use: "sub",
	Short: "subtract two numbers",
	Long: "subtract two numbers sub <number1> <number2>",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := operation("sub", args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("result: %f\n", result)
	},
	Args: cobra.ExactArgs(2),
}

var cmdMul = &cobra.Command{
	Use: "mul",
	Short: "multiply two numbers",
	Long: "multiply two numbers mul <number1> <number2>",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := operation("mul", args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("result: %f\n", result)
	},
	Args: cobra.ExactArgs(2),
}

var cmdDiv = &cobra.Command{
	Use: "div",
	Short: "divide two numbers",
	Long: "divide two numbers div <number1> <number2>",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := operation("div", args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("result: %f\n", result)
	},
	Args: cobra.ExactArgs(2),
}

func main() {
	var calculator = &cobra.Command{
		Use: "calculator",
		Short: "a simple calculator",
	}

	calculator.AddCommand(cmdAdd)
	calculator.AddCommand(cmdSub)
	calculator.AddCommand(cmdMul)
	calculator.AddCommand(cmdDiv)
	calculator.Execute()
}