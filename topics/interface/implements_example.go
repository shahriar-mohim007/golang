import (
	"math",
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
    rectangle := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 2}

    PrintArea(rectangle) 
    PrintArea(circle)    
}


Consider the following interface:

type Copier interface {
  Copy(string, string) int
}

// Based on the code alone, can you deduce what kinds of strings you should pass into the Copy function?

// We know the function signature expects 2 string types, but what are they? Filenames? URLs? Raw string data? For that matter, what the heck is that int that's being returned?

// Let's add some named arguments and return data to make it more clear.

// type Copier interface {
//   Copy(sourceFile string, destinationFile string) (bytesCopied int)
// }

// Much better. We can see what the expectations are now. The first argument is the sourceFile, 
// the second argument is the destinationFile, and bytesCopied, an integer, is returned.