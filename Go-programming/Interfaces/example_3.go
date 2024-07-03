package Interfaces

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

type Value interface {
	String() string
	Set(string) error
}

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) String() string {
	return fmt.Sprintf("%g°C", f.Celsius)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	_, err := fmt.Sscanf(s, "%f%s", &value, &unit)
	if err != nil {
		return err
	}
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
