package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// avoid goroutine leak, use default select or timeout
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case r := <-ch:
			fmt.Println("Got value", r)
		case <-time.After(time.Second * 4):
		default:
			fmt.Println("No value, moving on")
		}
	}()

	wg.Wait()

}
