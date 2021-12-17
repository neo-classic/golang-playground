package main

import (
	"fmt"
)

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: not found", e.Name)
}

func main() {
	err := getError()

	if e, ok := err.(*NotFoundError); ok {
		fmt.Println(e.Error())
		return
	}

	fmt.Println("no errors")
}

func getError() error {
	return &NotFoundError{
		Name: "The page",
	}
}
