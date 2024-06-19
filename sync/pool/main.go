package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			return "new item"
		},
	}

	item := pool.Get().(string)
	fmt.Println(item)

	pool.Put("used item")

	item = pool.Get().(string)
	fmt.Println(item)
}
