package main

import "fmt"

func main() {
	a := make(chan bool, 10)
	b := make(chan bool, 10)
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			fmt.Println("< a")
		case <-b:
			fmt.Println("< b")
		case <-c:
			fmt.Println("< c")
		default:
			fmt.Println("< default")
		}
	}
}
