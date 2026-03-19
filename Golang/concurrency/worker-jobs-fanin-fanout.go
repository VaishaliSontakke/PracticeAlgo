package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs { // workers compete to get jobs if not available block goroutine
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // signals no more jobs present for goroutines.

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}
