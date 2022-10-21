package main

import "fmt"

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
