// Эхо-сервер
package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("failed to open port on 1026")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection")
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("fatal error: %s", err)
		}
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("failed to read from socket")
	}

	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	conn.Write(data)
	panic(errors.New("pretend I'm a real error"))
}