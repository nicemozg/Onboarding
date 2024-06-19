package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialized")
}

func main() {
	for i := 0; i < 3; i++ {
		go once.Do(initialize)
	}
}
