package main

import "fmt"

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hola Anonymous"
	}

	return fmt.Sprintf("Hola %s", name)
}
