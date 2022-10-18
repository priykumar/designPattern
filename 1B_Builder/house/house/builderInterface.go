package main

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
