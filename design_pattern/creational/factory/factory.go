//Purpose: Defines an interface for creating objects but allows subclasses to alter the type of objects that will be created.

//Use Case: When a class cannot anticipate the type of objects it needs to create, or when the creation logic is complex and should be separated from the main class.

//The phrase "lets subclasses alter the type of objects that will be created" means that the Factory Method Pattern allows subclasses to control what specific type of object will be instantiated when the factory method is called. In other words, the parent class (or interface) defines the general mechanism for creating objects, but the subclasses provide the actual implementation that determines which specific class or object will be created.

package main

import "fmt"

// Product Interface (Vehicle)
type Vehicle interface {
	Drive() string
}

// Concrete Products (Car and Bike)
type Car struct{}
type Bike struct{}

// Implement the Drive method for Car
func (c Car) Drive() string {
	return "Driving a car!"
}

// Implement the Drive method for Bike
func (b Bike) Drive() string {
	return "Riding a bike!"
}

// Factory Method (Vehicle Factory)
func GetVehicle(vehicleType string) Vehicle {
	if vehicleType == "car" {
		return Car{} // Return Car object
	} else if vehicleType == "bike" {
		return Bike{} // Return Bike object
	}
	return nil
}

func main() {
	// Use the factory method to create different vehicles
	car := GetVehicle("car")
	bike := GetVehicle("bike")

	// Call the methods on the created objects
	fmt.Println(car.Drive())  // Output: Driving a car!
	fmt.Println(bike.Drive()) // Output: Riding a bike!
}
