package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Goroutine", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished")
}
