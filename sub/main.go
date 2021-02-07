package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// root context
var ctx = context.Background()

func main() {
	fmt.Println("init sub")

	rclient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pubsub := rclient.Subscribe(ctx, "chan1")
	fmt.Println("subscribed to chan1")

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println("==> receiving msg")
		fmt.Println(msg.Channel, msg.Payload)
	}
}
