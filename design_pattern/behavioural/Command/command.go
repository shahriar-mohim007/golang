package main

import "fmt"

// Command Interface
type Command interface {
	Execute()
}

// Receiver
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

// Concrete Commands
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(cmd Command) {
	r.command = cmd
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}
	remote := &RemoteControl{}

	onCommand := &LightOnCommand{light}
	offCommand := &LightOffCommand{light}

	remote.SetCommand(onCommand)
	remote.PressButton()

	remote.SetCommand(offCommand)
	remote.PressButton()
}
