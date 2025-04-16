package main

import "fmt"

type Strategy interface {
	Execute(a, b int) int
}

type Add struct{}

func (Add) Execute(a, b int) int { return a + b }

type Multiply struct{}

func (Multiply) Execute(a, b int) int { return a * b } 

func main() {
	var strategy Strategy = Add{}
	fmt.Println("Add:", strategy.Execute(2, 3))

	strategy = Multiply{}
	fmt.Println("Multiply:", strategy.Execute(2, 3))
}
