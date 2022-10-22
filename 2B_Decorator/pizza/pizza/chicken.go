package main

type chickenTopping struct {
	pizza pizzaInterface
}

func (c *chickenTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 100
}
func (c *chickenTopping) getConstituents() string {
	pizzaMaterial := c.pizza.getConstituents() + " + chicken"
	return pizzaMaterial
}
