package main

import (
	"fmt"
	"sync"
)

func workerDo(id int, jobs <-chan int, results chan<- int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			results <- j * 2
		case <-done:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	done := make(chan struct{})

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go workerDo(w, jobs, results, done, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // signals no more jobs present for goroutines.

	done <- struct{}{}

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}
