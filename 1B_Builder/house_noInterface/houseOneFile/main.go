package main

import "fmt"

// main struct: House
// builder struct: iglooHouseBuilder, normalHouseBuilder
// motive is to create object of main struct, thru builder structs

type House struct {
	window string
	door   string
	floor  string
	wall   string
}

type iglooHouseBuilder struct {
	window string
	door   string
	floor  string
	wall   string
}

type normalHouseBuilder struct {
	window string
	door   string
	floor  string
	wall   string
}

func newIglooBuilder() *iglooHouseBuilder {
	return &iglooHouseBuilder{}
}

func (b *iglooHouseBuilder) buildWindow() {
	b.window = "snowy window"
}

func (b *iglooHouseBuilder) buildDoor() {
	b.door = "snowy door"
}

func (b *iglooHouseBuilder) buildFloor() {
	b.floor = "snowy floor"
}

func (b *iglooHouseBuilder) buildWall() {
	b.wall = "rock wall"
}

func (b *iglooHouseBuilder) getHouse() House {
	return House{
		window: b.window,
		wall:   b.wall,
		floor:  b.floor,
		door:   b.door,
	}
}

func newNormalBuilder() *normalHouseBuilder {
	return &normalHouseBuilder{}
}

func (b *normalHouseBuilder) buildWindow() {
	b.window = "wooden window"
}

func (b *normalHouseBuilder) buildDoor() {
	b.door = "wooden door"
}

func (b *normalHouseBuilder) buildFloor() {
	b.floor = "concrete floor"
}

func (b *normalHouseBuilder) buildWall() {
	b.wall = "concrete wall"
}

func (b *normalHouseBuilder) getHouse() House {
	return House{
		window: b.window,
		wall:   b.wall,
		floor:  b.floor,
		door:   b.door,
	}
}

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

// Points to enhance:
// 	1.	Building doors, floors, windows, walls are basic neccesities of a house. Hence we can club
// 		buildDoor(), buildFloor(), buildWindow() and buildWall() inside an interface{}

// 	2.	Everytime a new house need to be created, we need to explictly call buildDoor(), buildFloor(),
// 		buildWindow() and buildWall(). Hence can create a new method that sequentially calls these methods, but that new
// 		method has to be present in new struct. That new struct is call Director
