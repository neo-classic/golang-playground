package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more arguments")
		os.Exit(1)
	}

	arguments := os.Args
	var sum int64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseInt(arguments[i], 10, 64)
		if err == nil {
			sum += n
			fmt.Println(n)
		}
	}
	fmt.Println("Sum is: ", sum)
}
