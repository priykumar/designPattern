package main

type chocolateChip struct {
	icecream icecreamInterface
}

func (c *chocolateChip) getPrice() int {
	price := c.icecream.getPrice()
	return price + 30
}
func (c *chocolateChip) getConstituents() string {
	return c.icecream.getConstituents() + " + chocolate chip"
}
