package main

import "fmt"

type dellLaptop struct {
	model string
}

func (dell *dellLaptop) connectToHdmi() {
	fmt.Printf("HDMI connector is plugged into %v\n\n", dell.model)
}

func (dell *dellLaptop) connectToTypeA() {
	fmt.Printf("Type-A connector is plugged into %v\n\n", dell.model)
}
