package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	go func() {
		mu.Lock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Goroutine notified")
		mu.Unlock()
	}()

	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
}
