package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(3.1415627))
	fmt.Println(strings.Title("head first go"))
	s := "Hello"

	for _, v := range s {
		fmt.Println(v)
	}
}
