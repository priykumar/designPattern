package main

type chocolateCone struct {
}

func (c *chocolateCone) getPrice() int {
	return 100
}

func (c *chocolateCone) getConstituents() string {
	return "chocolate cone"
}
