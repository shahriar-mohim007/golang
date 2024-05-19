// While Go is not object-oriented, it does support methods that can be defined on structs.
//  Methods are just functions that have a receiver. A receiver is a special parameter that
//  syntactically goes before the name of the function.

// type rect struct {
//   width int
//   height int
// }

// // area has a receiver of (r rect)
// func (r rect) area() int {
//   return r.width * r.height
// }

// var r = rect{
//   width: 5,
//   height: 10,
// }

// fmt.Println(r.area())
// // prints 50

// A receiver is just a special kind of function parameter. Receivers are important because they will,
// as you'll learn in the exercises to come, allow us to define interfaces that our structs (and other types) can implement.

package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
}

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
