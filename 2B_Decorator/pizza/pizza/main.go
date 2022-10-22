package main

import "fmt"

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
