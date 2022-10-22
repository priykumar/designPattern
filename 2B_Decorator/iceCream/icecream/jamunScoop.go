package main

type jamunScoop struct {
	icecream icecreamInterface
}

func (c *jamunScoop) getPrice() int {
	price := c.icecream.getPrice()
	return price + 80
}
func (c *jamunScoop) getConstituents() string {
	return c.icecream.getConstituents() + " + jamun scoop"
}
