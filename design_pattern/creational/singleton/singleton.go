//Purpose: Ensures that a class has only one instance and provides a global point of access to it.

//Use Case: When only one instance of a class is needed throughout the application, such as in logging, database connections, or thread pools.

package main

import (
	"fmt"
	"sync"
)

// Singleton structure
type Singleton struct{}

var instance *Singleton
var once sync.Once

// GetInstance returns the single instance of the Singleton class
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	// Both s1 and s2 will point to the same instance
	fmt.Println(s1 == s2) // Output: true
}
