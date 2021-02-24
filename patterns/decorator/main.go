package main

import (
	"fmt"
	"github.com/neo-classic/golang-playground/patterns/decorator/decorators"
	pizza2 "github.com/neo-classic/golang-playground/patterns/decorator/pizza"
)

func main() {
	pizza := &pizza2.VeggieMania{}

	pizzaWithCheese := &decorators.CheeseTopping{
		Pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &decorators.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price for veggieMania with tomato and cheese topping is: %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
