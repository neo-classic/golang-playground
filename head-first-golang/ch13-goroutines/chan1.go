package main

import "fmt"

func hiChan(ch chan string) {
	for i := 0; i < 4; i++ {
		ch <- fmt.Sprintf("hello channel #%d", i)
	}
}

func main() {
	ch := make(chan string)

	go hiChan(ch)

	t := <-ch
	fmt.Println(t)

	t = <-ch
	fmt.Println(t)
}
