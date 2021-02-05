package main

import "fmt"

const pi = 3.141
const (
	hello = "Привет"
	e     = 2.718
)
const (
	zero = iota
	_    = 2 * iota // пустая переменная, пропуск iota
	two
	three // = 3
)
const (
	_         = iota             // пропускаем первое значение
	KB uint64 = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	MB                           // 1 << (10 * 2) = 1048576
)
const (
	// нетипизированная константа
	year = 2017
	// типизированная константа
	yearTyped int = 2017
)

const (
	c1 = iota
	c2 = 2*iota + 1
	c3
	c4
	c5
)

func main() {
	var month int32 = 13
	fmt.Println(month + year)

	fmt.Println("zero, two, three: ", zero, two, three)

	fmt.Println("c1 c2 c3 c4 c5", c1, c2, c3, c4, c5)
	// month + yearTyped (mismatched types int32 and int)
	// fmt.Println( month + yearTyped )
}
