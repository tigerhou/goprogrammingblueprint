package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func main() {
	count := make(chan int)
	wg1.Add(2)
	fmt.Println("Starting Goroutines")
	go printCount("A", count)
	go printCount("B", count)
	fmt.Println("Channel begin")
	count <- 1
	fmt.Println("Waiting to Finish")
	wg1.Wait()
	fmt.Println("\nTerminating Program")

}

func printCount(label string, count chan int) {
	defer wg1.Done()
	for {
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}

		fmt.Printf("Count:%d received from %s\n", val, label)

		if val == 10 {
			fmt.Printf("Channel closed from %s\n ", label)
			close(count)
			return
		}
		val++
		count <- val
	}
}
