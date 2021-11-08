package main

import (
	"fmt"
	"time"
)

func main() {
	tickCall := time.NewTicker(2 * time.Second)
	defer tickCall.Stop()

	for {
		select {
		case <-tickCall.C:
			fmt.Println("qwe")
		default:
		}
		fmt.Println("ewq")
	}
}
