package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Membuat koneksi ke RabbitMQ
	conn, err := amqp.Dial("amqp://guest:10.50.12.236:55982/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Membuat channel untuk komunikasi
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	fmt.Println("Connected to RabbitMQ successfully!")

	// Mendeklarasikan queue
	// q, err := ch.QueueDeclare(
	// 	"hello", // Nama queue
	// 	true,    // Durable
	// 	false,   // Auto-delete
	// 	false,   // Exclusive
	// 	false,   // No-wait
	// 	nil,     // Arguments
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to declare a queue: %s", err)
	// }

	// // Mengirim pesan ke RabbitMQ
	// body := "Hello from Go!"
	// err = ch.Publish(
	// 	"",     // Default exchange
	// 	q.Name, // Routing key (queue)
	// 	false,  // Mandatory
	// 	false,  // Immediate
	// 	amqp.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        []byte(body),
	// 	})
	// if err != nil {
	// 	log.Fatalf("Failed to publish a message: %s", err)
	// }
	// log.Printf("Sent: %s", body)
}
