package main

import "fmt"

type pizzaInterface interface {
	getPrice() int
	getConstituents() string
}

type Pizza struct {
}

func (v *Pizza) getPrice() int {
	return 150
}
func (v *Pizza) getConstituents() string {
	pizzaMaterial := "pizza bread"
	return pizzaMaterial
}

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

func main() {
	// pizza1 := &Pizza{}
	// pizzaWithCheese := &cheeseTopping{
	// 	pizza: pizza1,
	// }
	// pizzaWithCheeseAndTomato := &tomatoTopping{
	// 	pizza: pizzaWithCheese,
	// }

	vegPizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: &cheeseTopping{
			pizza: &Pizza{},
		},
	}

	nonVegPizza := &chickenTopping{
		pizza: &tomatoTopping{
			pizza: &cheeseTopping{
				pizza: &Pizza{},
			},
		},
	}

	// fmt.Println(pizzaWithCheeseAndTomato.getPrice())
	fmt.Println(vegPizzaWithCheeseAndTomato.getPrice())
	fmt.Println(vegPizzaWithCheeseAndTomato.getConstituents())

	fmt.Println(nonVegPizza.getPrice())
	fmt.Println(nonVegPizza.getConstituents())
}
