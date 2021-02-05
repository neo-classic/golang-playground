package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars result"
}

func testFunc() {
	defer func() {
		fmt.Println("defer testFunc")
	}()
	fmt.Println("qwe")
}

func main() {
	defer fmt.Println("After work")
	defer func() {
		fmt.Println(getSomeVars())
	}()
	testFunc()
	fmt.Println("Some userful work")
}
