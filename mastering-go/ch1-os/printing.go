package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"

	fmt.Println("====fmt.Print====")
	fmt.Print(v1, v2, v3, v4)
	fmt.Println()

	fmt.Println("====fmt.Println====")
	fmt.Println(v1, v2, v3, v4)
	fmt.Println()

	fmt.Println("====fmt.Print & space====")
	fmt.Print(v1, " ", v2, " ", v3, " ", v4)
	fmt.Println()

	fmt.Println("====fmt.Printf====")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)

	fmt.Println("====Verbs====")
	fmt.Printf("T: %T %T\n", v1, v2)
	fmt.Printf("05d: %05d\n", v2)
}
