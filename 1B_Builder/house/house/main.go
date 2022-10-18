package main

import "fmt"

func main() {
	// create builder object
	normalHouseBuilder := getBuilder("Normal")

	// create a new director and associate a builder with it
	newDirector := newDirector()
	newDirector.setBuilder(normalHouseBuilder)

	// build a house thru director
	newNormalHouse := newDirector.buildHouse()
	fmt.Printf("\n%#v\n\n", newNormalHouse)

	iglooHouseBuilder := getBuilder("Igloo")
	newDirector.setBuilder(iglooHouseBuilder)
	newIglooHouse := newDirector.buildHouse()
	fmt.Printf("%#v\n\n", newIglooHouse)
}
