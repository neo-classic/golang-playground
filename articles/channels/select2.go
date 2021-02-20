package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan bool, 10)
	b := make(chan bool, 10)

	go func() {
		time.Sleep(time.Second * 10)
		for i := 0; i < 10; i++ {
			a <- true
			b <- true
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			fmt.Print("< a ")
		case <-b:
			fmt.Print("< b ")
		}
	}
}
