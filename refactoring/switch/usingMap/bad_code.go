package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	someDay := Monday

	switch someDay {
	case Monday:
		one()
	case Tuesday:
		two()
	case Wednesday:
		three()
	default:
		fmt.Println("we don't support this day")
	}
}

func one() {
	fmt.Println("one")
}

func two() {
	fmt.Println("two")
}

func three() {
	fmt.Println("three")
}
