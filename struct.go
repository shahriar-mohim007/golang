// // Maps are a convenient way to store some kinds of data, but they have
// // limitations. They don’t define an API since there’s no way to constrain a
// // map to only allow certain keys. Also, all of the values in a map must be of the same type. For these reasons, maps are not an ideal way to pass data
// // from function to function. When you have related data that you want to group together, you should define a struct.
package main

//
//import "fmt"
//
//// You can define a
//// struct type inside or outside of a function. A struct type that’s defined
//// within a function can only be used within that function.
//type person struct {
//	name string
//	age int
//	pet string
//}
//func main(){
//	var fred person
//	bob := person{}
//	julia := person{
//		"Julia",
//		40,
//		"cat",
//	}
//	beth := person{
//		age: 30,
//		name: "Beth",
//	}
//	bob.name = "Bob"
//    fmt.Println(beth.name)
//	//Anonymous Structs
//	var person struct {
//		name string
//		age int
//		pet string
//	}
//	person.name = "bob"
//	person.age = 50
//	person.pet = "dog"
//	pet := struct {
//	name string
//	kind string
//	}{
//	name: "Fido",
//	kind: "dog",
//	}
//
////There are two common situations where
//// anonymous structs are handy. The first is when you translate external data
//// into a struct or a struct into external data (like JSON or protocol buffers).
//// This is called marshaling and unmarshaling data. We’ll learn how to do this
//// in “encoding/json”.
////Go does allow you to perform a type
//// conversion from one struct type to another if the fields of both structs have
//// the same names, order, and types.
//type firstPerson struct {
//	name string
//	age int
//}
//type secondPerson struct {
//	name string
//	age int
//}
//
//type firstPerson struct {
//	name string
//	age int
//}
//	f := firstPerson{
//	name: "Bob",
//	age: 50,
//}
//	var g struct {
//	name string
//	age int
//}
//	// compiles -- can use = and == between identical named and anonymous structs
//	g = f
//	fmt.Println(f == g)
//
//
//// 	If all the fields of a struct are comparable, the struct itself is comparable, so two expressions of
//// that type may be compared using == or !=. The == operation compares the corresponding
//// fields of the two structs in order, so the two printed expressions below are equivalent:
//type Point struct{ X, Y int }
//p := Point{1, 2}
//q := Point{2, 1}
//fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
//fmt.Println(p == q)
//// "false"
//// Comparable struct types, like other comparable types, may be used as the key type of a map.
//type address struct {
//hostname string
//port int
//}
//hits := make(map[address]int)
//hits[address{"golang.org", 443}]++
//}
