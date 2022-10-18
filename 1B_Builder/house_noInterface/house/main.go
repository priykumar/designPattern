package main

import "fmt"

func main() {
	// create a builder object
	normalHouse := newNormalBuilder()

	// set all attributes
	normalHouse.buildDoor()
	normalHouse.buildFloor()
	normalHouse.buildWindow()
	normalHouse.buildWall()

	// create object of main struct using object of a builder struct
	newNormalHouse := normalHouse.getHouse()
	fmt.Printf("\n%#v\n\n", newNormalHouse)

	iglooHouse := newIglooBuilder()
	iglooHouse.buildDoor()
	iglooHouse.buildFloor()
	iglooHouse.buildWall()
	iglooHouse.buildWindow()
	newIglooHouse := iglooHouse.getHouse()
	fmt.Printf("%#v\n\n", newIglooHouse)

}
