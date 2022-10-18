package main

import "fmt"

// main struct: House
// builder struct: iglooHouseBuilder, normalHouseBuilder
// motive is to create object of main struct, thru builder structs

type builderInterface interface {
	buildWindow()
	buildDoor()
	buildFloor()
	buildWall()
	getHouse() House
}

func getBuilder(builderType string) builderInterface {
	if builderType == "Normal" {
		return newNormalBuilder()
	} else if builderType == "Igloo" {
		return newIglooBuilder()
	}
	return nil
}

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

type Director struct {
	builder builderInterface
}

func newDirector() *Director {
	return &Director{}
}

func (d *Director) setBuilder(b builderInterface) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.buildDoor()
	d.builder.buildFloor()
	d.builder.buildWall()
	d.builder.buildWindow()
	return d.builder.getHouse()
}

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

// Points to enhance:
// 	1.	Building doors, floors, windows, walls are basic neccesities of a house. Hence we can club
// 		buildDoor(), buildFloor(), buildWindow() and buildWall() inside an interface{}

// 	2.	Everytime a new house need to be created, we need to explictly call buildDoor(), buildFloor(),
// 		buildWindow() and buildWall(). Hence can create a new method that sequentially calls these methods, but that new
// 		method has to be present in new struct. That new struct is call Director
