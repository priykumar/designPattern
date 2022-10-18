package main

type normalHouseBuilder struct {
	window string
	door   string
	floor  string
	wall   string
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
