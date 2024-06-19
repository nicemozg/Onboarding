package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	var wg sync.WaitGroup
	const numWorkers = 3

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(in, out, &wg)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	for result := range out {
		fmt.Println("Result:", result)
	}
}

func worker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		out <- n * n
	}
}
