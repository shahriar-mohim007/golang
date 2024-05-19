// An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.

// if INITIAL_STATEMENT; CONDITION {
// }

// Why would I use this?

// This is just some syntactic sugar that Go offers to shorten up code in some cases. For example, instead of writing:

// length := getLength(email)
// if length < 1 {
//     fmt.Println("Email is invalid")
// }

// We can do:

// if length := getLength(email); length < 1 {
//     fmt.Println("Email is invalid")
// }
