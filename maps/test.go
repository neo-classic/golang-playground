package main

import (
	"fmt"
	"time"
)

func main() {
	mp1 := make(map[string]int, 1)
	fmt.Printf("%+v, len = %d, address = %p\n", mp1, len(mp1), mp1)

	mp1["a"] = 1
	fmt.Printf("%+v, len = %d, address = %p\n", mp1, len(mp1), mp1)

	mp1["b"] = 2
	fmt.Printf("%+v, len = %d, address = %p\n", mp1, len(mp1), mp1)

	mp1["c"] = 3
	fmt.Printf("%+v, len = %d, address = %p\n", mp1, len(mp1), mp1)

	mp1["d"] = 4
	fmt.Printf("%+v, len = %d, address = %p\n", mp1, len(mp1), mp1)

	if val, ok := mp1["Qwe"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("Qwe key doesn't exists in the map")
	}

	delete(mp1, "d")
	fmt.Printf("%+v, len = %d, address = %p\n", mp1, len(mp1), mp1)

	qwe := map[string]time.Time{}
	t, _ := qwe["www"]
	t.Add(123 * time.Hour)
	fmt.Println(t)
	fmt.Println(time.Now().Before(t))
}
