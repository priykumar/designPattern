package main

type chocolateScoop struct {
	icecream icecreamInterface
}

func (c *chocolateScoop) getPrice() int {
	price := c.icecream.getPrice()
	return price + 100
}
func (c *chocolateScoop) getConstituents() string {
	return c.icecream.getConstituents() + " + chocolate scoop"
}
