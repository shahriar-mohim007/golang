//package main
//
//import (
//	"fmt"
//)
//
//func f(from string) {
//	for i := 0; i < 3; i++ {
//		fmt.Println(from, ":", i)
//	}
//}
//
//func main() {
//
//	f("direct")
//
//	go f("goroutine")
//
//	go func(msg string) {
//		fmt.Println(msg)
//	}("going")
//
//	//time.Sleep(time.Minute)
//	fmt.Println("done")
//}

//In your Go program, the time.Sleep(time.Minute)
//function is crucial for observing the output of the goroutines. Here's why:
//
//Goroutines and Concurrency: Goroutines are concurrent
//threads of execution. When you call go f("goroutine")
//and go func(msg string) { fmt.Println(msg) }("going"),
//these functions run asynchronously, meaning the
// main function doesn't wait for them to complete.
//
// Main Function Exit: If the main function
// completes before the goroutines,
// the program will exit, and the goroutines
// won't have a chance to run.
//The time.Sleep(time.Minute) call ensures that the main
//function doesn't exit immediately, giving the goroutines time to execute.

//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//	// Get the number of CPU cores that Go is using
//	cores := runtime.GOMAXPROCS(0) // Passing 0 just returns the current value without changing it
//	fmt.Printf("Go is using %d cores\n", cores)
//}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("Factorial of %d is %d\n", n, result)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Utilize all available CPU cores
	var wg sync.WaitGroup

	numbers := []int{20, 7, 9, 11, 12, 3, 10, 5}

	for _, num := range numbers {
		wg.Add(1)
		go factorial(num, &wg) // Run factorial computations in parallel
	}

	wg.Wait() // Wait for all goroutines to finish
}
