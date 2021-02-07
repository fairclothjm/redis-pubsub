package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// root context
var ctx = context.Background()

func main() {
	fmt.Println("init pub")

	rclient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		fmt.Printf("<== published msg %d\n", i)
		err := rclient.Publish(ctx, "chan1", fmt.Sprintf("payload-%d", i)).Err()
		if err != nil {
			panic(err)
		}
	}
}
