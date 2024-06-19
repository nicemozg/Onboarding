package main

import (
	"context"
	"fmt"
)

type key int

const myKey key = 0

func main() {
	ctx := context.WithValue(context.Background(), myKey, "myValue")

	val := ctx.Value(myKey)
	fmt.Println("Value associated with key:", val)
}
