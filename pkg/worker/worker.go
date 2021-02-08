package worker

import (
	"log"
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

// Worker attempts to provision a new resource for each job and
// writes the results of the privision call to the results chan.
func Worker(id int, jobs <-chan int, results chan<- Result) {
	for j := range jobs {
		log.Println("worker", id, "started job", j)
		resp := provisionNewResource()

		provisioned := false
		if resp == http.StatusOK {
			provisioned = true
		}

		results <- Result{ID: j, Provisioned: provisioned}
		log.Println("worker", id, "finished job", j)
	}
}
