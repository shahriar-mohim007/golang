package main

import "fmt"

// Component interface
type Coffee interface {
	Cost() int
}

// Concrete Component
type BasicCoffee struct{}

func (c *BasicCoffee) Cost() int {
	return 5 // Basic coffee costs $5
}

// Decorator
type CoffeeDecorator struct {
	coffee Coffee
}

func (c *CoffeeDecorator) Cost() int {
	return c.coffee.Cost()
}

// Concrete Decorators
type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(c Coffee) Coffee {
	return &MilkDecorator{CoffeeDecorator{coffee: c}}
}

func (m *MilkDecorator) Cost() int {
	return m.coffee.Cost() + 2 // Adding milk costs $2
}

type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(c Coffee) Coffee {
	return &SugarDecorator{CoffeeDecorator{coffee: c}}
}

func (s *SugarDecorator) Cost() int {
	return s.coffee.Cost() + 1 // Adding sugar costs $1
}

// Client code
func main() {
	basicCoffee := &BasicCoffee{}
	fmt.Printf("Cost of basic coffee: $%d\n", basicCoffee.Cost())

	// Adding milk
	coffeeWithMilk := NewMilkDecorator(basicCoffee)
	fmt.Printf("Cost of coffee with milk: $%d\n", coffeeWithMilk.Cost())

	// Adding milk and sugar
	coffeeWithMilkAndSugar := NewSugarDecorator(coffeeWithMilk)
	fmt.Printf("Cost of coffee with milk and sugar: $%d\n", coffeeWithMilkAndSugar.Cost())
}
