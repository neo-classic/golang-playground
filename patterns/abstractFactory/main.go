package main

import (
	"fmt"
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory"
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shirt"
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shoe"
)

func main() {
	adidasFactory, _ := factory.GetSportsFactory(factory.BrandAdidas)
	nikeFactory, _ := factory.GetSportsFactory(factory.BrandNike)

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s shoe.ShoeInterface) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShirtDetails(s shirt.ShirtInterface) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
