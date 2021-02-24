package decorators

import "github.com/neo-classic/golang-playground/patterns/decorator/pizza"

// конкретный декратор
type TomatoTopping struct {
	Pizza pizza.Pizza
}

func (t *TomatoTopping) GetPrice() int {
	pizzaPrice := t.Pizza.GetPrice()
	return pizzaPrice + 7
}
