package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Printf("b: address=%p, value=%d\n", b, *b)
	*b = 3  // a = 3
	c := &a // новый указатель на переменную a
	*c = 123
	fmt.Printf("a-addr=%p, b-addr=%p, c-addr=%p, c-val=%d\n", &a, b, c, *c)

	// получение указателя на переменнут типа int
	// инициализировано значением по-умолчанию
	d := new(int)
	fmt.Println(d)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c и a не изменились

	c = d   // теперь с указывает туда же, куда d
	*c = 14 // с = 14 -> d = 14, a = 12
}
