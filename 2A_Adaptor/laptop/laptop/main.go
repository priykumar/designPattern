package main

func main() {
	client := &client{}
	dell := &dellLaptop{
		model: "Inspiron",
	}
	mac := &macBook{
		model: "Macbook air",
	}
	adptr := &macAdaptor{
		mac: mac,
	}

	client.connectToLaptopPort(dell)
	client.connectToLaptopPort(adptr)
}
