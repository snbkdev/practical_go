// Закрытие канала на стороне отправителя
package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	
	go send(ch)

	for {
		select {
		case m, ok := <- ch:
			if !ok {
				log.Println("channel closed")
				return
			}
			log.Println("Got message: ", m)
		case <- timeout:
			log.Println("time out")
			return
		default:
			log.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	log.Println("sent and closed")
}