package main

import (
	"fmt"
	"runtime"
)

func printString(s string, times int) {
	for i := 0; i < times; i++ {
		fmt.Print(s)
		runtime.Gosched()
	}
}

func doSome(ch chan string, str string) {
	ch <- str
}

func main() {
	strs := []string{
		"a",
		"b",
		"c",
		"d",
	}

	ch := make(chan string)

	for _, s := range strs {
		go doSome(ch, s)
	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//
	//for _, s := range strs {
	//	go printString(s, 10)
	//}
	//
	//time.Sleep(1 * time.Second)
	//
	//fmt.Println("")

}
