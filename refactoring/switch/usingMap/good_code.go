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
	someDay := Tuesday

	var dayMap = map[Weekday]func(){
		Monday:    one,
		Tuesday:   two,
		Wednesday: three,
	}

	if dayFunc, ok := dayMap[someDay]; ok {
		dayFunc()
	} else {
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
