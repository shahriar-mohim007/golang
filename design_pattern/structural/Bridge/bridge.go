package main

import "fmt"

// Implementor
type Color interface {
	ApplyColor()
}

type RedColor struct{}

func (rc *RedColor) ApplyColor() { fmt.Println("Applying red color") }

type BlueColor struct{}

func (bc *BlueColor) ApplyColor() { fmt.Println("Applying blue color") }

// Abstraction
type Shape interface {
	Draw()
}

type Circle struct {
	color Color
}

func (c *Circle) Draw() {
	c.color.ApplyColor()
	fmt.Println("Drawing Circle")
}

func main() {
	red := &RedColor{}
	blue := &BlueColor{}

	circle1 := &Circle{color: red}
	circle2 := &Circle{color: blue}

	circle1.Draw()
	circle2.Draw()
}

// package main

// import "fmt"

// // Implementor interface (Device)
// type Device interface {
//     PowerOn()
//     PowerOff()
// }

// // Concrete Implementations (TV and Radio)
// type TV struct{}

// func (t *TV) PowerOn() {
//     fmt.Println("TV is now ON")
// }

// func (t *TV) PowerOff() {
//     fmt.Println("TV is now OFF")
// }

// type Radio struct{}

// func (r *Radio) PowerOn() {
//     fmt.Println("Radio is now ON")
// }

// func (r *Radio) PowerOff() {
//     fmt.Println("Radio is now OFF")
// }

// // Abstraction (RemoteControl)
// type RemoteControl struct {
//     device Device
// }

// func (r *RemoteControl) TurnOn() {
//     r.device.PowerOn()
// }

// func (r *RemoteControl) TurnOff() {
//     r.device.PowerOff()
// }

// // Client code
// func main() {
//     tv := &TV{}
//     radio := &Radio{}

//     tvRemote := &RemoteControl{device: tv}
//     tvRemote.TurnOn()  // TV is now ON
//     tvRemote.TurnOff() // TV is now OFF

//     radioRemote := &RemoteControl{device: radio}
//     radioRemote.TurnOn()  // Radio is now ON
//     radioRemote.TurnOff() // Radio is now OFF
// }
