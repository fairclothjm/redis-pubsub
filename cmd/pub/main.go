package main

import (
	"context"
	"encoding/json"
	"fmt"
	w "github.com/fairclothjm/redis-pubsub/pkg/worker"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

// root context
var ctx = context.Background()

func main() {
	log.Println("init pub")

	rclient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	const jobCount = 20
	jobs := make(chan int, jobCount)
	results := make(chan w.Result, jobCount)

	// start workers
	const workerCount = 3
	for i := 1; i <= workerCount; i++ {
		go w.Worker(i, jobs, results)
	}

	// send the jobs and close channel
	for j := 1; j <= jobCount; j++ {
		jobs <- j
	}

	// collect results
	// alternatively, use waitgroup?
	for i := 1; i <= jobCount; i++ {
		r := <-results
		log.Printf("%+v\n", r)

		if r.Provisioned {
			marshalValue, err := json.Marshal(r)
			if err != nil {
				panic(err)
			}

			err = rclient.Set(ctx, strconv.Itoa(r.ID), marshalValue, 0).Err()
			if err != nil {
				panic(err)
			}
			log.Printf("add provisioned resource with ID %d to cache", r.ID)
		} else {
			log.Printf("failure for %d, send alert", r.ID)
			err := rclient.Publish(ctx, "chan_provision_failure", fmt.Sprintf("payload-%d", r.ID)).Err()
			if err != nil {
				panic(err)
			}
		}
	}
}
