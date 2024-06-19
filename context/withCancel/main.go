package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancel() // отменяем контекст через 2 секунды
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context is done due to cancellation")
	}
}
