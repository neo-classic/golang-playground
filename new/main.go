package main

import "fmt"

type t1 struct {
	str string
	v   int
}

func main() {
	obj1 := new(t1)
	obj1.str = "qwe"
	var obj2 t1
	obj3 := &t1{str: "ewq"}

	fmt.Printf("obj1 = %+v\n", obj1) // obj1 = &{str:qwe v:0}
	fmt.Printf("obj2 = %+v\n", obj2) // obj2 = {str: v:0}
	fmt.Printf("obj3 = %+v\n", obj3) // obj3 = &{str:ewq v:0}
	fmt.Println("------------------")

	sl1 := make([]int, 10)
	sl1[3] = 3
	sl2 := []int{}
	sl2 = append(sl2, 4)

	mp1 := make(map[string]int, 10)
	mp2 := map[string]int{}
	mp2["sss"] = 5

	fmt.Printf("sl1 = %+v, type = %T\n", sl1, sl1)
	fmt.Printf("sl2 = %+v, type = %T\n", sl2, sl2)
	fmt.Printf("mp1 = %+v, type = %T\n", mp1, mp1)
	fmt.Printf("mp2 = %+v, type = %T\n", mp2, mp2)
}
