package main

import "fmt"

func main() {
	qwe := make(map[string]bool)
	qwe["qwe1"] = true
	qwe["qwe2"] = true

	for q := range qwe {
		fmt.Println(q)
	}
}
