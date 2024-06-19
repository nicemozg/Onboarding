package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("key", "value")

	value, ok := m.Load("key")
	if ok {
		fmt.Println("Loaded value:", value)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
