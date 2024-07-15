package main

import (
	"fmt"
	"sync"
)

// Global variables
var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func increment() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock() // Lock the mutex before accessing the shared resource
		counter++
		mutex.Unlock() // Unlock the mutex after accessing the shared resource
	}
}

func main() {
	wg.Add(2)      //adds two goroutines to the wait group.
	go increment() // Start first goroutine
	go increment() // Start second goroutine

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
