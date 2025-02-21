package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Starting Peril server...")
	const conn_string = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(conn_string)
    if err != nil {
        fmt.Printf("Cannot connect to %v. %w\n", conn_string, err)
		return
    }
	defer conn.Close()
	fmt.Printf("Connected to %v. \n", conn_string)
	ch, err := conn.Channel()
    if err != nil {
        fmt.Printf("Cannot get channel. %w\n", err)
    }

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Shutting down...")
}
