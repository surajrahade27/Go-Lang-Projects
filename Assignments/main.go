package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type ResponseData struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
}

var jobs = make(chan Job, 100)
var results = make(chan ResponseData, 100)

const Url = "https://xkcd.com"

func fetch(n int) (*ResponseData, error) {

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	// concatenate strings to get url; ex: https://xkcd.com/571/info.0.json
	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	var data ResponseData

	// error from web service, empty struct to avoid disruption of process
	if resp.StatusCode != http.StatusOK {
		data = ResponseData{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}

	resp.Body.Close()
	return &data, nil
}

type Job struct {
	number int
}

func allocateJobs(noOfJobs int) {
	for i := 0; i <= noOfJobs; i++ {
		jobs <- Job{i + 1}
	}
	close(jobs)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		result, err := fetch(job.number)
		if err != nil {
			log.Printf("error in fetching: %v\n", err)
		}
		results <- *result
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i <= noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func getResults(done chan bool) {
	for result := range results {
		if result.Num != 0 {
			fmt.Printf("=================\n")
			fmt.Printf("Retrieving No : #%d\n", result.Num)
			fmt.Printf("Safe Title : #%s\n", result.SafeTitle)
		}
	}
	done <- true
}

func main() {
	// allocate jobs
	noOfJobs := 10
	go allocateJobs(noOfJobs)

	// get results
	done := make(chan bool)
	go getResults(done)

	// create worker pool
	noOfWorkers := 3
	createWorkerPool(noOfWorkers)
	//worker -> fetch
	// wait for all results to be collected
	<-done

}
