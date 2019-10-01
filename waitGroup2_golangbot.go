/*
We will implement worker pool using buffered channels. Our worker pool will carry out the task of finding the sum of a digits of the input number. For example if 234 is passed, the output would be 9 (2 + 3 + 4). The input to the worker pool will be list of pseudo random integers.
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	jobs    = make(chan Job, 10)
	results = make(chan Result, 10)

	noOfWorkers  = flag.Int("workers", 1, "No of Worker routines")
	noOfJobs     = flag.Int("jobs", 1, "No of Jobs")
	printResults = flag.Bool("print", false, "True if you want to print results")
)

// Each Job struct has id and randomNo for the sum of individual digits to be computed
type Job struct {
	id       int
	randomNo int
}

// The Results struct has a job filed which is the job for which it holds the results (sum of individual digits) in the sumOfDigits field
type Result struct {
	job         Job
	sumOfDigits int
}

// The digits fn below does teh actual job of finding the sum of individual digits of an integer and returning it.
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// Get work from jobs chan and do the work
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomNo)}
		results <- output
	}
	wg.Done()
}

// Create worker goroutines based on noOfWorker
func createWorkerPool(noOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// To populate the jobs chan
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{id: i, randomNo: randomNo}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		if *printResults {
			fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomNo, result.sumOfDigits)
		}
	}
	done <- true
}

func main() {
	flag.Parse()
	startTime := time.Now()
	noOfJobs := *noOfJobs
	noOfWorkers := *noOfWorkers
	go allocate(noOfJobs) // populates job's channel
	done := make(chan bool)
	go createWorkerPool(noOfWorkers) // will create go routines
	go result(done)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken ", diff.Seconds(), "seconds")

}
