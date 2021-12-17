package main

import "fmt"

func main() {
	sl := make([]int, 2, 2)
	fmt.Printf("%+v, len = %d, cap = %d\n", sl, len(sl), cap(sl))

	sl = append(sl, 1)
	fmt.Printf("%+v, len = %d, cap = %d\n", sl, len(sl), cap(sl))

	sl = append(sl, 2, 3)
	fmt.Printf("%+v, len = %d, cap = %d\n", sl, len(sl), cap(sl))

	sl2 := []int{}
	fmt.Printf("%+v, len = %d, cap = %d\n", sl2, len(sl2), cap(sl2))

	sl2 = append(sl2, 1)
	fmt.Printf("%+v, len = %d, cap = %d\n", sl2, len(sl2), cap(sl2))

	sl2 = append(sl2, 2)
	fmt.Printf("%+v, len = %d, cap = %d\n", sl2, len(sl2), cap(sl2))

	sl2 = append(sl2, 3)
	fmt.Printf("%+v, len = %d, cap = %d\n", sl2, len(sl2), cap(sl2))

	sl2 = append(sl2, []int{4, 5}...)
	fmt.Printf("%+v, len = %d, cap = %d\n", sl2, len(sl2), cap(sl2))
}
