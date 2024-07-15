package main

import (
	"fmt"
	"sync"
	"time"
)

// worker is a function that simulates doing some work
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work with a sleep
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	const numWorkers = 5
	wg.Add(numWorkers) //Set the number of goroutines to wait for

	for i := 1; i <= numWorkers; i++ {
		fmt.Println("wg is:", &wg)
		go worker(i, &wg) // Start a worker goroutine
	}

	wg.Wait() // Wait for all worker goroutines to complete
	fmt.Println("All workers done...")
}
