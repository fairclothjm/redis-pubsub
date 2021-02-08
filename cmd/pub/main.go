package main

import (
	"context"
	"fmt"
	w "github.com/fairclothjm/redis-pubsub/pkg/worker"
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

	const jobCount = 20
	jobs := make(chan int, jobCount)
	results := make(chan Result, jobCount)

	// start 5 workers
	for i := 1; i <= 3; i++ {
		go w.worker(i, jobs, results)
	}

	// send the jobs and close channel
	for j := 1; j <= jobCount; j++ {
		jobs <- j
	}

	// collect results
	for i := 1; i <= jobCount; i++ {
		r := <-results
		fmt.Printf("%+v\n", r)
		if !r.provisioned {
			fmt.Println(" ==> failure, send alert")
			err := rclient.Publish(ctx, "chan1", fmt.Sprintf("payload-%d", r.id)).Err()
			if err != nil {
				panic(err)
			}
		}
	}
}
