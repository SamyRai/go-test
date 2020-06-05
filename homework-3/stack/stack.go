package stack

import (
	"fmt"
	"github.com/SamyRai/go-test/homework-3/auto"
)

var QueueToBuyCar []auto.Car

func AddCar(car auto.Car) {
	QueueToBuyCar = append(QueueToBuyCar, car)
}

func PopCar(){
	QueueToBuyCar = QueueToBuyCar[1:]
}

func PrintInfo(){
	fmt.Println("Currently we have", len(QueueToBuyCar), "orders")
}

func ShowAll(){
	for i, car := range QueueToBuyCar {
		fmt.Println(i+1, "-", car.Brand.Name, car.Brand.Model, "left")
	}
}
