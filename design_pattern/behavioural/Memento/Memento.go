package main

import "fmt"

// Memento holds the state of the text
type Memento struct {
	state string
}

// GetState returns the saved state
func (m *Memento) GetState() string {
	return m.state
}

// TextEditor is the originator that we want to save/restore the state of
type TextEditor struct {
	content string
}

// Write updates the content of the text editor
func (e *TextEditor) Write(text string) {
	e.content += text
}

// Save creates a new memento with the current state
func (e *TextEditor) Save() *Memento {
	return &Memento{state: e.content}
}

// Restore restores the state from a memento
func (e *TextEditor) Restore(memento *Memento) {
	e.content = memento.GetState()
}

// GetContent returns the current content of the text editor
func (e *TextEditor) GetContent() string {
	return e.content
}

// Caretaker manages mementos (history of states)
type Caretaker struct {
	history []*Memento
}

// AddMemento saves a memento to the history
func (c *Caretaker) AddMemento(m *Memento) {
	c.history = append(c.history, m)
}

// GetMemento retrieves a memento from the history
func (c *Caretaker) GetMemento(index int) *Memento {
	if index < len(c.history) {
		return c.history[index]
	}
	return nil
}

func main() {
	// Create the originator (text editor) and the caretaker (history manager)
	editor := &TextEditor{}
	caretaker := &Caretaker{}

	// Write some text and save states
	editor.Write("Hello, ")
	caretaker.AddMemento(editor.Save())

	editor.Write("World!")
	caretaker.AddMemento(editor.Save())

	editor.Write(" How are you?")
	fmt.Println("Current Content:", editor.GetContent()) // Hello, World! How are you?

	// Undo the last change
	editor.Restore(caretaker.GetMemento(1))
	fmt.Println("After Undo:", editor.GetContent()) // Hello, World!

	// Undo again
	editor.Restore(caretaker.GetMemento(0))
	fmt.Println("After Second Undo:", editor.GetContent()) // Hello,
}
