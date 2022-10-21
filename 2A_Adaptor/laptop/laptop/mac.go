package main

import "fmt"

type macBook struct {
	model string
}

func (m *macBook) connectToTypeC() {
	fmt.Printf("Type-C connector is plugged into %v\n", m.model)
}
