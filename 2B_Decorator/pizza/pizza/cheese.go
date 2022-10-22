package main

type cheeseTopping struct {
	pizza pizzaInterface
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 70
}
func (c *cheeseTopping) getConstituents() string {
	pizzaMaterial := c.pizza.getConstituents() + " + cheese"
	return pizzaMaterial
}
