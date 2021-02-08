package worker

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Result struct {
	ID          int
	Provisioned bool
}

var mockResponses = []int{
	http.StatusOK,
	http.StatusInternalServerError,
}

// doMockAPICall "randomly" returns a 200 or 500 response
func doMockAPICall() int {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(mockResponses))
	time.Sleep(time.Second)

	return mockResponses[n]
}

func provisionNewResource() int {
	return doMockAPICall()
}

func Worker(id int, jobs <-chan int, results chan<- Result) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		resp := provisionNewResource()

		provisioned := false
		if resp == http.StatusOK {
			provisioned = true
		}

		results <- Result{id: j, provisioned: provisioned}
		fmt.Println("worker", id, "finished job", j)
	}
}
