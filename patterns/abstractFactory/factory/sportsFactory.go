package factory

import (
	"fmt"
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shirt"
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shoe"
)

const (
	BrandAdidas = "adidasFactory"
	BrandNike   = "nike"
)

type SportsFactoryInterface interface {
	MakeShoe() shoe.ShoeInterface
	MakeShirt() shirt.ShirtInterface
}

func GetSportsFactory(brand string) (SportsFactoryInterface, error) {
	switch brand {
	case BrandAdidas:
		return &adidasFactory{}, nil
	case BrandNike:
		return &nikeFactory{}, nil
	default:
		return nil, fmt.Errorf("Do not support the brand")
	}
}
