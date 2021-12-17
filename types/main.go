package main

import "fmt"

func main() {
	v1 := 1
	v2 := "qwe"
	v3 := &v1

	fmt.Printf("v1 has type = %T and value = %d\n", v1, v1)
	fmt.Printf("v2 has type = %T and value = %s\n", v2, v2)
	fmt.Printf("v3 has type = %T and address = %p, value = %d\n", v3, v3, *v3)

	var t interface{}
	t = v1
	switch t := t.(type) {
	default:
		fmt.Printf("t has type = %T\n", t)
	case int:
		fmt.Printf("t has type = integer and val = %d\n", t)
	}
}
