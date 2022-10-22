package main

type sprinkler struct {
	icecream icecreamInterface
}

func (c *sprinkler) getPrice() int {
	price := c.icecream.getPrice()
	return price + 20
}
func (c *sprinkler) getConstituents() string {
	return c.icecream.getConstituents() + " + sprinkler"
}
