package main

import "fmt"

type Car struct {
	Wheels int
	Color  string
}

type CarBuilder struct {
	car Car
}

func (c *CarBuilder) AddWheels(count int) *CarBuilder {
	c.car.Wheels = count
	return c
}

func (c *CarBuilder) AddColor(color string) *CarBuilder {
	c.car.Color = color
	return c
}

func (cb *CarBuilder) GetCar() Car {
	return cb.car
}

func main() {
	carBuilder := &CarBuilder{}
	car := carBuilder.AddWheels(4).AddColor("red").GetCar()
	fmt.Println(car)
}
