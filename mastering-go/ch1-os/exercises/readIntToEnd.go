package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "END" {
			return
		} else {
			fmt.Println(">>> ", input)
		}
	}
}
