package main

import (
	"fmt"
	"time"
)

// Function to simulate some work and send a result through a channel
func worker(id int, ch chan<- string) {
	time.Sleep(time.Second)                 // Simulate work with a sleep
	ch <- fmt.Sprintf("Worker %d done", id) // Send result to channel
}

func main() {
	// Create a channel to communicate between goroutines
	ch := make(chan string)

	// Start a few worker goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receive and print results from the channel
	for i := 1; i <= 3; i++ {
		result := <-ch
		fmt.Println(result)
	}
}
