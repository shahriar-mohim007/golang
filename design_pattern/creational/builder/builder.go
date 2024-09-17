//The Builder Pattern is a creational design pattern that allows you to construct complex objects step by step. It separates the construction of an object from its representation, allowing you to
//create different representations of the same object by using the same construction process.

//When to Use:

// When the object to be built is complex with many optional fields.
// When you want to create different variations of an object using the same building process.
// When you want to ensure the immutability of an object after it's been built.

package main

import "fmt"

// Product: Computer
type Computer struct {
	CPU     string
	GPU     string
	RAM     string
	Storage string
}

func (c Computer) String() string {
	return fmt.Sprintf("Computer with CPU: %s, GPU: %s, RAM: %s, Storage: %s", c.CPU, c.GPU, c.RAM, c.Storage)
}

// Builder Interface
type ComputerBuilder interface {
	SetCPU() ComputerBuilder
	SetGPU() ComputerBuilder
	SetRAM() ComputerBuilder
	SetStorage() ComputerBuilder
	GetComputer() Computer
}

// Concrete Builder: Gaming Computer
type GamingComputerBuilder struct {
	computer Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{}
}

func (b *GamingComputerBuilder) SetCPU() ComputerBuilder {
	b.computer.CPU = "High-Performance CPU"
	return b
}

func (b *GamingComputerBuilder) SetGPU() ComputerBuilder {
	b.computer.GPU = "High-End Gaming GPU"
	return b
}

func (b *GamingComputerBuilder) SetRAM() ComputerBuilder {
	b.computer.RAM = "32GB DDR4"
	return b
}

func (b *GamingComputerBuilder) SetStorage() ComputerBuilder {
	b.computer.Storage = "1TB SSD"
	return b
}

func (b *GamingComputerBuilder) GetComputer() Computer {
	return b.computer
}

// Concrete Builder: Workstation Computer
type WorkstationComputerBuilder struct {
	computer Computer
}

func NewWorkstationComputerBuilder() *WorkstationComputerBuilder {
	return &WorkstationComputerBuilder{}
}

func (b *WorkstationComputerBuilder) SetCPU() ComputerBuilder {
	b.computer.CPU = "Multi-Core Workstation CPU"
	return b
}

func (b *WorkstationComputerBuilder) SetGPU() ComputerBuilder {
	b.computer.GPU = "Workstation GPU"
	return b
}

func (b *WorkstationComputerBuilder) SetRAM() ComputerBuilder {
	b.computer.RAM = "64GB DDR4"
	return b
}

func (b *WorkstationComputerBuilder) SetStorage() ComputerBuilder {
	b.computer.Storage = "2TB SSD"
	return b
}

func (b *WorkstationComputerBuilder) GetComputer() Computer {
	return b.computer
}

// Director: Constructs the computer using the builder
type ComputerDirector struct {
	builder ComputerBuilder
}

func NewComputerDirector(builder ComputerBuilder) *ComputerDirector {
	return &ComputerDirector{
		builder: builder,
	}
}

func (d *ComputerDirector) BuildComputer() Computer {
	return d.builder.SetCPU().SetGPU().SetRAM().SetStorage().GetComputer()
}

func main() {
	// Build a Gaming Computer
	gamingBuilder := NewGamingComputerBuilder()
	director := NewComputerDirector(gamingBuilder)
	gamingComputer := director.BuildComputer()
	fmt.Println(gamingComputer) // Output: Computer with CPU: High-Performance CPU, GPU: High-End Gaming GPU, RAM: 32GB DDR4, Storage: 1TB SSD

	// Build a Workstation Computer
	workstationBuilder := NewWorkstationComputerBuilder()
	director = NewComputerDirector(workstationBuilder)
	workstationComputer := director.BuildComputer()
	fmt.Println(workstationComputer) // Output: Computer with CPU: Multi-Core Workstation CPU, GPU: Workstation GPU, RAM: 64GB DDR4, Storage: 2TB SSD
}
