package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // важно вызывать cancel, чтобы освободить ресурсы

	select {
	case <-ctx.Done():
		fmt.Println("Context is done due to deadline")
	}
}
