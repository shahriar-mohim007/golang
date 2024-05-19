// Go is not an object-oriented language. However, embedded structs provide a kind of data-only
// inheritance that can be useful at times. Keep in mind, Go doesn't
// support classes or inheritance in the complete sense, but embedded structs are a way to elevate and share fields between struct definitions.

// type car struct {
//   maker string
//   model string
// }

// type truck struct {
//   // "car" is embedded, so the definition of a
//   // "truck" now also additionally contains all
//   // of the fields of the car struct
//   car
//   bedSize int
// }

// Embedded vs nested

//     Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
//     Like nested structs, you assign the promoted fields with the embedded struct in a composite literal.

// lanesTruck := truck{
//   bedSize: 10,
//   car: car{
//     maker: "toyota",
//     model: "camry",
//   },
// }

// fmt.Println(lanesTruck.bedSize)

// // embedded fields promoted to the top-level
// // instead of lanesTruck.car.maker
// fmt.Println(lanesTruck.maker)
// fmt.Println(lanesTruck.model)

package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

// don't edit below this line

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
