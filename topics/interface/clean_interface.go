import (
	"io"
	"os"
)

//Clean Interfaces

// Writing clean interfaces is hard. Frankly, anytime you’re dealing with abstractions
// in code, the simple can become complex very quickly if you’re not careful. Let’s go over some rules of thumb for keeping interfaces clean.
// 1. Keep Interfaces Small

// If there is only one piece of advice that you take away from this article, make it this: keep interfaces small!
// Interfaces are meant to define the minimal behavior necessary to accurately represent an idea or concept.

// Here is an example from the standard HTTP package of a larger interface that’s a good example of defining minimal behavior:

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

// Interfaces Should Have No Knowledge of Satisfying Types

// An interface should define what is necessary
// for other types to classify as a member of that interface. They shouldn’t be aware of any types that happen to satisfy the interface at design time.

// For example, let’s assume we are building an interface to describe the components necessary to define a car.

type car interface {
	Color() string
	Speed() int
	IsFiretruck() bool
}

//Which inherits the required methods from car and adds one additional required method to make the car a firetruck.

// Interfaces are not classes, they are slimmer.
// Interfaces don’t have constructors or deconstructors that require that data is created or destroyed.
// Interfaces aren’t hierarchical by nature, though there is syntactic sugar to create interfaces that happen to be supersets of other interfaces.
// Interfaces define function signatures, but not underlying behavior. Making an interface often won’t DRY up your code in regards to struct methods.
// For example, if five types satisfy the fmt.Stringer interface, they all need their own version of the String() function.
