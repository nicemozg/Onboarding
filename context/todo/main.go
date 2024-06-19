package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()

	select {
	case <-ctx.Done():
		fmt.Println("Context is done")
	default:
		fmt.Println("Context is not done")
	}
}
