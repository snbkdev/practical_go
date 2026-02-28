// Пример микросервиса
package main

import (
	"context"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := "media"
	kafkaHost := os.Getenv("KAFKA_HOST")
	if kafkaHost == "" {
		panic("KAFKA_HOST environment variable not set")
	}

	conn, err := kafka.DialLeader(
		context.Background(), "tcp", kafkaHost, topic, 0, // Подключается к экземпляру Kafka по протоколу TCP
	)
	if err != nil {
		panic(err)
	}

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6)	 // Создает пакет для чтения сообщений с указанием допустимого диапазона размеров
	message := make([]byte, 10e3)

	for {
		n, err := batch.Read(message) 	// Обрабатывает(читает) сообщения 
		if err != nil {
			break
		}
		log.Println(string(message[:n]))
	}

	batch.Close()
	conn.Close()
}