package main 

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	fmt.Println("Hello World!")
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {

	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to connect open a Channel")
	q, err := ch.QueueDeclare("hello", false /*durable*/, false /*autoDelete*/, false, false, nil) 
	
}

func failOnError(err Error, msg string) {

	if err != nil {
		log.Fatalf("%s: %s", msg, err)	
		panic(fmt.Sprintf("%s: %s", msg, err))
	}	


}