package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i //Produce a value
		//time.Sleep(time.Millisecond * 500)
	}
	close(ch) //Close the channel when done producing
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Consumed:", value) //Consume a value
	}
}

func main() {
	ch := make(chan int)

	go producer(ch) //Start producer goroutine
	go consumer(ch) //Start consumer goroutine

	time.Sleep(time.Second * 4) //Wait for goroutine to finish
}
