package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return fmt.Sprint("this is dog")
}

type Cat struct{}

func (c Cat) Speak() string {
	return fmt.Sprint("this is cat")
}

func GetAnimal(animal string) Animal {
	switch animal {
	case "dog":
		return Dog{}
	case "cat":
		return Cat{}
	default:
		return nil
	}
}

func main() {
	animal := GetAnimal("dog")
	fmt.Println(animal.Speak())
	animal = GetAnimal("cat")
	fmt.Println(animal.Speak())
}
