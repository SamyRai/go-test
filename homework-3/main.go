package main

import (
	"github.com/SamyRai/go-test/homework-3/auto"
	"github.com/SamyRai/go-test/homework-3/stack"
)

func main() {
	honda := auto.Car{
		Brand: &auto.Brand {
			Name:  "Honda",
			Model: "Civic",
		},
		Tech: &auto.Tech {
				YearOfProduction: "2012",
				HorsePower:       134,
				MaxSpeed:         180,
		},
		MaxPassengers: 5,
		Color: "red",

	}

	toyota := auto.Car{
		Brand: &auto.Brand {
			Name:  "Toyota",
			Model: "Corolla",
		},
		Tech: &auto.Tech {
			YearOfProduction: "2018",
			HorsePower:       154,
			MaxSpeed:         190,
		},
		MaxPassengers: 5,
		Color: "yellow",

	}

	kamaz := auto.Truck{
		Brand: &auto.Brand {
			Name:  "Kamaz",
			Model: "Z500",
		},
		Tech: &auto.Tech {
			YearOfProduction: "1990",
			HorsePower:       600,
			MaxSpeed:         110,
		},
		MaxWeight: 3000,
		Height: 600,
	}

	honda.ShowSellingInfo()
	kamaz.ShowSellingInfo()
	toyota.ShowSellingInfo()

	honda.Tech.YearOfProduction = "2020" // little cheating on production date

	honda.ShowSellingInfo()

	// Adding cars to the queue for buyers

	stack.AddCar(honda)
	stack.AddCar(toyota)
	stack.PrintInfo()

	stack.PopCar()
	stack.PrintInfo()
	stack.ShowAll() // Honda should be gone

}

