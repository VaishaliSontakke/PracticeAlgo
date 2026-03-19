package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// Next send blocks until a value is received
	go func() { ch <- 3 }()
	fmt.Println(<-ch)

}
