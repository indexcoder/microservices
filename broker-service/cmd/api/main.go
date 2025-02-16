package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

const webPort = "8080"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {

	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println("this is 1 = ", err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	app := Config{
		Rabbit: rabbitConn,
	}
	log.Printf("Starting server on port: %s", webPort)
	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var connection *amqp.Connection
	var backoff time.Duration

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Printf("RabbitMQ not yet ready broker servce... %v\n", err)
			counts++
		} else {
			log.Printf("RabbitMQ connected... broker servce ...")
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
