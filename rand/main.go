package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("1:", rand.Int())
	fmt.Println("2:", rand.Int())
	fmt.Println("3:", rand.Int())

	fmt.Println(RandomString(10))
	fmt.Println(RandomString(10))
	fmt.Println(RandomString(10))

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int())
	}
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
