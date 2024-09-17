package main

import (
	"fmt"
)

// Flyweight interface
type Character interface {
	Display(font string, size int, color string)
}

// Concrete Flyweight
type ConcreteCharacter struct {
	intrinsicState string // Shared state
}

func (c *ConcreteCharacter) Display(font string, size int, color string) {
	fmt.Printf("Character: %s, Font: %s, Size: %d, Color: %s\n",
		c.intrinsicState, font, size, color)
}

// Flyweight Factory
type CharacterFactory struct {
	characters map[string]*ConcreteCharacter
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{characters: make(map[string]*ConcreteCharacter)}
}

func (f *CharacterFactory) GetCharacter(char string) *ConcreteCharacter {
	if _, exists := f.characters[char]; !exists {
		f.characters[char] = &ConcreteCharacter{intrinsicState: char}
	}
	return f.characters[char]
}

// Client
func main() {
	factory := NewCharacterFactory()

	chars := []struct {
		char  string
		font  string
		size  int
		color string
	}{
		{"A", "Arial", 12, "Red"},
		{"B", "Arial", 12, "Blue"},
		{"A", "Times New Roman", 14, "Green"},
	}

	for _, c := range chars {
		character := factory.GetCharacter(c.char)
		character.Display(c.font, c.size, c.color)
	}
}
