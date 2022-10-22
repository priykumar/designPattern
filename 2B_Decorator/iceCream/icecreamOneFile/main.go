package main

import "fmt"

type icecreamInterface interface {
	getPrice() int
	getConstituents() string
}

type chocolateCone struct {
}

func (c *chocolateCone) getPrice() int {
	return 100
}

func (c *chocolateCone) getConstituents() string {
	return "chocolate cone"
}

type vanillaCone struct {
}

func (v *vanillaCone) getPrice() int {
	return 50
}

func (v *vanillaCone) getConstituents() string {
	return "vanilla cone"
}

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

func main() {
	icecream1 := &chocolateChip{
		icecream: &sprinkler{
			icecream: &chocolateScoop{
				icecream: &chocolateCone{},
			},
		},
	}

	fmt.Println(icecream1.getPrice())
	fmt.Println(icecream1.getConstituents())

	icecream2 := &sprinkler{
		icecream: &chocolateChip{
			icecream: &jamunScoop{
				icecream: &vanillaCone{},
			},
		},
	}
	fmt.Println(icecream2.getPrice())
	fmt.Println(icecream2.getConstituents())
}
