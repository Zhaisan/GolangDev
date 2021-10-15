package main

import (
    "fmt"
    "time"
)

const cntJobs = 8

func main() {

    jobs := make(chan int, cntJobs)
    results := make(chan int, cntJobs)

	for i := 1; i <= 7; i++ {
        go worker(i, jobs, results)
    }

	for j := 1; j <= cntJobs; j++ {
        jobs <- j
    }
	close(jobs)

	for k := 1; k <= cntJobs; k++ {
        <-results
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {

    for x := range jobs {
        fmt.Printf("Worker %v started job %v\n", id, x)
        time.Sleep(time.Second)
        fmt.Printf("Worker %v finished job %v\n", id, x)
        results <- x * 2
    }
}