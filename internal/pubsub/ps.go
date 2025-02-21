package pubsub

import (
	"encoding/json"
	"fmt"
	"sync/context"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	jbytes, err := json.Marshal(val)
    if err != nil {
        fmt.Printf("Cannot read from response. %w", err)
		return
    }
	ch.PublishWithContext(context.Background(), exchange, key, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: jbytes,
	})
}


[]
