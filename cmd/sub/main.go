package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

// root context
var ctx = context.Background()

func main() {
	log.Println("init sub")

	rclient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pubsub := rclient.Subscribe(ctx, "chan1")
	log.Println("subscribed to chan1")

	ch := pubsub.Channel()
	for msg := range ch {
		log.Println("==> receiving msg", msg.Channel, msg.Payload)
	}
}
