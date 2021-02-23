package factory

import (
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shirt"
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shoe"
)

type adidasFactory struct {
}

func (a *adidasFactory) MakeShoe() shoe.ShoeInterface {
	return &shoe.AdidasShoe{
		Shoe: shoe.Shoe{
			Logo: BrandAdidas,
			Size: 14,
		},
	}
}

func (a *adidasFactory) MakeShirt() shirt.ShirtInterface {
	return &shirt.AdidasShirt{
		Shirt: shirt.Shirt{
			Logo: BrandAdidas,
			Size: 15,
		},
	}
}
