package main

import (
	"fmt"
	"sync"
)

func main() {

	// pass arg i to avoid closure bug
	// The common "closure bug" in Go occurs when a closure launched as a goroutine
	//or used in a deferred function references a loop variable which is updated
	//in each iteration, but not redeclared. This results in all closures capturing
	//the reference to the single variable, and thus using its final value after
	//the loop finishes, rather than the value at the time the closure was created.

	// in short if variable is not redeclared the go routine picks final value of loop.
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println("Hello World", val)
		}(i)
	}
	wg.Wait()

}
