package auto

import "fmt"

type Tech struct {
	YearOfProduction string
	HorsePower uint
	MaxSpeed uint
}

type Brand struct {
	Name string
	Model string
}

type Car struct {
	Tech *Tech
	Brand *Brand

	MaxPassengers uint
	Color string
}

func (car Car) ShowSellingInfo() {
	fmt.Printf(
		"Hello, you are looking at %s %s %s which was produced at %s and has increadable %d horse powers and %d max speed\n",
		car.Color,
		car.Brand.Name,
		car.Brand.Model,
		car.Tech.YearOfProduction,
		car.Tech.HorsePower,
		car.Tech.MaxSpeed)
}

type Truck struct {
	Tech *Tech
	Brand *Brand

	MaxWeight uint
	Height uint
}

func (truck Truck) ShowSellingInfo() {
	fmt.Printf(
		"Hello, you are looking at powerfull truck %s %s which was produced at %s and has increadable %d horse powers and %d max speed and can carry %d kg\n",
		truck.Brand.Name,
		truck.Brand.Model,
		truck.Tech.YearOfProduction,
		truck.Tech.HorsePower,
		truck.Tech.MaxSpeed,
		truck.MaxWeight)
}
