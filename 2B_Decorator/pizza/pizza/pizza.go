package main

type Pizza struct {
}

func (v *Pizza) getPrice() int {
	return 150
}
func (v *Pizza) getConstituents() string {
	pizzaMaterial := "pizza bread"
	return pizzaMaterial
}
