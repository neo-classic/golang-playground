package main

import "fmt"

func main() {
	lfu := &lfu{}
	cache := NewCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	lru := &lru{}
	cache.setEvictionAlgo(lru)

	cache.add("c", "3")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("d", "4")

	cache.get("d")

	fmt.Println(cache.storage)
}
