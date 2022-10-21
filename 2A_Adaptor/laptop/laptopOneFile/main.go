package main

import "fmt"

type client struct {
}

func (c *client) connectToLaptopPort(laptop connectToPorts) {
	laptop.connectToHdmi()
	laptop.connectToTypeA()
}

type connectToPorts interface {
	connectToHdmi()
	connectToTypeA()
}

type dellLaptop struct {
	model string
}

func (dell *dellLaptop) connectToHdmi() {
	fmt.Printf("HDMI connector is plugged into %v\n\n", dell.model)
}

func (dell *dellLaptop) connectToTypeA() {
	fmt.Printf("Type-A connector is plugged into %v\n\n", dell.model)
}

type macBook struct {
	model string
}

func (m *macBook) connectToTypeC() {
	fmt.Printf("Type-C connector is plugged into %v\n", m.model)
}

type macAdaptor struct {
	mac *macBook
}

func (a *macAdaptor) connectToHdmi() {
	a.mac.connectToTypeC()
	fmt.Printf("     HDMI connector is plugged into %v via mac adaptor\n\n", a.mac.model)
}

func (a *macAdaptor) connectToTypeA() {
	a.mac.connectToTypeC()
	fmt.Printf("     Type-A connector is plugged into %v via mac adaptor\n\n", a.mac.model)
}

func main() {
	client := &client{}
	dell := &dellLaptop{
		model: "Inspiron",
	}
	mac := &macBook{
		model: "Macbook air",
	}
	adptr := &macAdaptor{
		mac: mac,
	}

	client.connectToLaptopPort(dell)
	client.connectToLaptopPort(adptr)
}
