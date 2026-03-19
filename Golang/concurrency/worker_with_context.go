package main

import (
	"context"
	"fmt"
	"sync"
)

/* Safe Cycle

start workers
send jobs
close jobs
wait workers
close results
consume results
// wait group Add → start goroutine → defer Done in routine  */

func workerCnt(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			results <- j * 2
		case <-ctx.Done():
			fmt.Printf("Worker %d received shutdown signal\n", id)
			return
		}
	}
}

func printResults(results <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range results {
		fmt.Println("Result:", r)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	jobs := make(chan int, 6)
	results := make(chan int, 6)

	for w := 0; w < 3; w++ {
		wg.Add(1)
		go workerCnt(ctx, w, jobs, results, &wg)
	}

	for i := 0; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}

}
