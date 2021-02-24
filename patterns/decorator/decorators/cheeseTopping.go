package decorators

import "github.com/neo-classic/golang-playground/patterns/decorator/pizza"

// конкретный декоратор
type CheeseTopping struct {
	Pizza pizza.Pizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
