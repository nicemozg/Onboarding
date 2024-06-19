package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // важно вызывать cancel, чтобы освободить ресурсы

	select {
	case <-ctx.Done():
		fmt.Println("Context is done due to timeout")
	}
}
