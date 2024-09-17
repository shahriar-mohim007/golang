package main

import "fmt"

// Subsystems
type Amplifier struct{}

func (a *Amplifier) On()  { fmt.Println("Amplifier on") }
func (a *Amplifier) Off() { fmt.Println("Amplifier off") }

type DVDPlayer struct{}

func (d *DVDPlayer) On()               { fmt.Println("DVD Player on") }
func (d *DVDPlayer) Play(movie string) { fmt.Println("Playing", movie) }

// Facade
type HomeTheaterFacade struct {
	amp *Amplifier
	dvd *DVDPlayer
}

func (htf *HomeTheaterFacade) WatchMovie(movie string) {
	htf.amp.On()
	htf.dvd.On()
	htf.dvd.Play(movie)
}

func main() {
	amp := &Amplifier{}
	dvd := &DVDPlayer{}
	facade := &HomeTheaterFacade{amp: amp, dvd: dvd}

	facade.WatchMovie("Inception")
}
