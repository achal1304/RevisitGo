package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

type Options func(*Product)

func NewProduct(options ...Options) *Product {
	p := &Product{}
	for _, opt := range options {
		opt(p)
	}
	return p
}

func WithName(name string) Options {
	return func(p *Product) {
		p.Name = name
	}
}

func WithPrice(price int) Options {
	return func(p *Product) {
		p.Price = price
	}
}

func main() {
	prod := NewProduct(WithName("test"), WithPrice(2))
	fmt.Println(prod)
}
