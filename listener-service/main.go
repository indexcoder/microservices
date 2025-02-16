package main

import (
	"fmt"
	ampq "github.com/rabbitmq/amqp091-go"
	"listener/event"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println("this is 1 = ", err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// start listening for message
	log.Println("Listening for and consuming RabbitMQ messages...")

	// create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Println("this is 2 = ", err)
	}

	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println("this is 3 = ", err)
	}
}

func connect() (*ampq.Connection, error) {
	var counts int64
	var connection *ampq.Connection
	var backoff time.Duration

	for {
		c, err := ampq.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Printf("RabbitMQ not yet ready... %v\n", err)
			counts++
		} else {
			log.Printf("RabbitMQ connected...")
			connection = c
			break
		}

		if counts >= 5 {
			fmt.Printf("ampq connection timed out\n")
			return nil, fmt.Errorf("connection attempts exceeded limit: %v", err)
		}

		backoff = time.Duration(math.Min(math.Pow(float64(counts), 2), 32)) * time.Second
		log.Printf("backing off ... %v\n", backoff)

		time.Sleep(backoff)
		continue
	}

	return connection, nil
}
