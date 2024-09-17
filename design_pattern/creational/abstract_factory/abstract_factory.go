//Purpose: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.

//Use Case: When a system needs to create objects that are part of a family of related products,
//but you want to abstract away from the specific product implementations.

package main

import "fmt"

// Abstract factory
type GUIFactory interface {
	CreateButton() Button
	CreateWindow() Window
}

// Abstract products
type Button interface {
	Click()
}

type Window interface {
	Open()
}

// Concrete products
type MacButton struct{}

func (b MacButton) Click() { fmt.Println("Mac button clicked") }

type MacWindow struct{}

func (w MacWindow) Open() { fmt.Println("Mac window opened") }

type WinButton struct{}

func (b WinButton) Click() { fmt.Println("Windows button clicked") }

type WinWindow struct{}

func (w WinWindow) Open() { fmt.Println("Windows window opened") }

// Concrete factories
type MacFactory struct{}

func (f MacFactory) CreateButton() Button { return MacButton{} }
func (f MacFactory) CreateWindow() Window { return MacWindow{} }

type WinFactory struct{}

func (f WinFactory) CreateButton() Button { return WinButton{} }
func (f WinFactory) CreateWindow() Window { return WinWindow{} }

func main() {
	var factory GUIFactory
	os := "mac" // Simulate determining the OS at runtime

	if os == "mac" {
		factory = MacFactory{}
	} else {
		factory = WinFactory{}
	}

	button := factory.CreateButton()
	window := factory.CreateWindow()

	button.Click() // Output: Mac button clicked
	window.Open()  // Output: Mac window opened
}
