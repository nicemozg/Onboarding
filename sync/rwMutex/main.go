package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	rwMutex sync.RWMutex
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		rwMutex.RLock()
		fmt.Println("Read Counter:", counter)
		rwMutex.RUnlock()
	}()

	go func() {
		defer wg.Done()
		rwMutex.Lock()
		for i := 0; i < 1000; i++ {
			counter++
		}
		rwMutex.Unlock()
	}()

	wg.Wait()
}
