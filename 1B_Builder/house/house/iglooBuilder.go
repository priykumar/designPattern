package main

type iglooHouseBuilder struct {
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
