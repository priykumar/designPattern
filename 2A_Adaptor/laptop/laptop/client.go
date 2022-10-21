package main

type client struct {
}

func (c *client) connectToLaptopPort(laptop connectToPorts) {
	laptop.connectToHdmi()
	laptop.connectToTypeA()
}
