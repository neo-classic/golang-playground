package main

import "fmt"

type Address interface {
	FullAddress() string
}

type Person struct {
	Name    string
	address string
}

func (p *Person) FullAddress() string {
	return p.address
}

type Office struct {
	Employers []Person
	address   string
}

func (o *Office) FullAddress() string {
	return o.address
}

func GetAddress(o Address) {
	fmt.Printf("%T\n", o)
	switch o.(type) {
	case *Person:
		fmt.Println("*Person")
	case *Office:
		fmt.Println("*Office")
	}
}

func main() {
	p := &Person{
		Name:    "Yuriy",
		address: "qwe123",
	}
	fmt.Println(p.FullAddress())
	GetAddress(p)
}
