package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var accessCount int

func incr() {
	mu.Lock() // this means that only single goroutine can has access it at a time!
	defer mu.Unlock()
	accessCount = accessCount + 1
}

func main() {
	var wg sync.WaitGroup
	accessCount = 0
	loop := 500
	wg.Add(loop) // amount of goroutines

	for i := 1; i <= loop; i++ {
		go func(c int) {
			defer wg.Done()
			incr()
		}(i)
	}

	wg.Wait()
	fmt.Println("Final count =", accessCount)
}
