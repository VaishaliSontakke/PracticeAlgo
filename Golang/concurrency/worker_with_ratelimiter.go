package main

import (
	"fmt"
	"sync"
	"time"
)

func workerRate(id int, jobs chan int, results chan int, wg *sync.WaitGroup, limiter <-chan time.Time) {
	defer wg.Done()
	for job := range jobs {
		<-limiter
		fmt.Println("worker", id, "start", job)
		results <- job * 2
		fmt.Println("worker", id, "Completed", job)

	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup
	// 2 ticks generated every second. two jobs will be started by worker every second.
	limiter := time.Tick(500 * time.Millisecond)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go workerRate(w, jobs, results, &wg, limiter)
	}

	for r := 1; r <= 5; r++ {
		jobs <- r
	}
	close(jobs)
	wg.Wait()
	close(results)
	for r := range results {
		fmt.Println("Result:", r)
	}
}
