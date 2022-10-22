package main

type tomatoTopping struct {
	pizza pizzaInterface
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 50
}
func (t *tomatoTopping) getConstituents() string {
	pizzaMaterial := t.pizza.getConstituents() + " + tomato"
	return pizzaMaterial
}
