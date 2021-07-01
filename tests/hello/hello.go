package main

import "fmt"

const helloPrefix = "Hello"

func Hello(msg string) string {
	if msg == "" {
		msg = "world"
	}
	return fmt.Sprintf("%s, %s", helloPrefix, msg)
}

func main() {
	fmt.Println(Hello("world"))
}
