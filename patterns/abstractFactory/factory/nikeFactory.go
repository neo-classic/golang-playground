package factory

import (
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shirt"
	"github.com/neo-classic/golang-playground/patterns/abstractFactory/factory/shoe"
)

type nikeFactory struct {
}

func (n *nikeFactory) MakeShoe() shoe.ShoeInterface {
	return &shoe.NikeShoe{
		Shoe: shoe.Shoe{
			Logo: BrandNike,
			Size: 9,
		},
	}
}

func (n *nikeFactory) MakeShirt() shirt.ShirtInterface {
	return &shirt.NikeShirt{
		Shirt: shirt.Shirt{
			Logo: BrandNike,
			Size: 8,
		},
	}
}
