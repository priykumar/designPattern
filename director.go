package main

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
