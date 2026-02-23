// Отправка трассировки стека в стандартный поток вывода
package main

import "runtime/debug"

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	debug.PrintStack()
}