// Простая блокировка с помощью каналов
package main

import (
	"log"
	"time"
)

func main() {
	lock := make(chan bool, 1) 	// создает канал с размером буфера6 равным 1 
	for i := 1; i < 7; i++ {
		go worker(i, lock) 	// запускает до 6 горутин, использующих общий канал
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	log.Printf("%d wants the lock\n", id)
	lock <- true	// одна из горутин захватывает блокировку, отправляя сообщение в канал. Остальные дожидаются своей очереди
	log.Printf("%d has the lock\n", id)		// участок кода между lock <- true и <-lock считается заблокированным этой горутиной
	<- lock		// снимает блокировку, считывая значение из канала. Это освобождает буфер и позволяет следующей горутине захватить блокировку
	log.Printf("%d is releasing the lock \n", id)
}