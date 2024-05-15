// An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.

// To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:

// myCar := struct {
//   maker string
//   model string
// } {
//   maker: "tesla",
//   model: "model 3",
// }

// You can even nest anonymous structs as fields within other structs:

// type car struct {
//   maker string
//   model string
//   doors int
//   mileage int
//   // wheel is a field containing an anonymous struct
//   wheel struct {
//     radius int
//     material string
//   }
// }
