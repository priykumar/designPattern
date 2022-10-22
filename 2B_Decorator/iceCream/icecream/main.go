package main

import "fmt"

func main() {
	icecream1 := &chocolateChip{
		icecream: &sprinkler{
			icecream: &chocolateScoop{
				icecream: &chocolateCone{},
			},
		},
	}

	fmt.Println(icecream1.getPrice())
	fmt.Println(icecream1.getConstituents())

	icecream2 := &sprinkler{
		icecream: &chocolateChip{
			icecream: &jamunScoop{
				icecream: &vanillaCone{},
			},
		},
	}
	fmt.Println(icecream2.getPrice())
	fmt.Println(icecream2.getConstituents())
}
