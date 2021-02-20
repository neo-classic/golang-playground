package main

import (
	"fmt"
	"github.com/neo-classic/golang-playground/patterns/singleton/single"
	"sync"
)

const iterCount = 10

func main() {
	var wg sync.WaitGroup

	wg.Add(iterCount)
	fmt.Println("Pure singleton realization:")
	for i := 0; i < iterCount; i++ {
		go single.GetInstance(&wg)
	}
	wg.Wait()
	fmt.Println("=====================")

	wg.Add(iterCount)
	fmt.Println("sync.Once realization:")
	for i := 0; i < iterCount; i++ {
		go single.GetInstanceTwo(&wg)
	}
	wg.Wait()
}
