package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide one or more args")
		os.Exit(1)
	}

	arguments := os.Args
	var sum float64
	var count float64

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			fmt.Println(n)
			sum += n
			count++
		}
	}

	fmt.Println("Sum: ", sum)
	fmt.Println("Avg: ", sum/count)
}
