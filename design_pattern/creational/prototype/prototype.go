//The Prototype Pattern is a creational design pattern that allows you to create new objects by copying an existing object,
//known as the prototype. Instead of creating objects from scratch, you clone or duplicate an existing object, which can then be modified as needed.
//This is useful when object creation is costly,
//or you want to avoid the complexities of creating an object from scratch.

package main

import "fmt"

// Prototype interface
type Shape interface {
	Clone() Shape
	GetInfo() string
}

// Concrete Prototype: Circle
type Circle struct {
	Radius int
}

func (c *Circle) Clone() Shape {
	// Return a new copy of Circle
	return &Circle{
		Radius: c.Radius,
	}
}

func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Circle with radius %d", c.Radius)
}

// Concrete Prototype: Rectangle
type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Clone() Shape {
	// Return a new copy of Rectangle
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
	}
}

func (r *Rectangle) GetInfo() string {
	return fmt.Sprintf("Rectangle with width %d and height %d", r.Width, r.Height)
}

func main() {
	// Create original shapes
	circle := &Circle{Radius: 10}
	rectangle := &Rectangle{Width: 20, Height: 15}

	// Clone shapes
	circleClone := circle.Clone()
	rectangleClone := rectangle.Clone()

	// Print information about original and cloned shapes
	fmt.Println(circle.GetInfo())      // Output: Circle with radius 10
	fmt.Println(circleClone.GetInfo()) // Output: Circle with radius 10

	fmt.Println(rectangle.GetInfo())      // Output: Rectangle with width 20 and height 15
	fmt.Println(rectangleClone.GetInfo()) // Output: Rectangle with width 20 and height 15
}
