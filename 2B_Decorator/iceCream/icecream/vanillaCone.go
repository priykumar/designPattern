package main

type vanillaCone struct {
}

func (v *vanillaCone) getPrice() int {
	return 50
}

func (v *vanillaCone) getConstituents() string {
	return "vanilla cone"
}
