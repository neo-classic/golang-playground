package main

import (
	"fmt"
	"runtime"
	"time"
)

func printString(s string, times int) {
	for i := 0; i < times; i++ {
		fmt.Print(s)
		runtime.Gosched()
	}
}

func main() {
	strs := []string{
		"a",
		"b",
		"c",
		"d",
	}

	for _, s := range strs {
		go printString(s, 10)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("")
}
