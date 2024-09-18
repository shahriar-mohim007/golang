package main

import (
    "fmt"
)

func divide(a, b int) int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    if b == 0 {
        panic("cannot divide by zero")
    }
    return a / b
}

func main() {
    fmt.Println(divide(10, 2)) // Works fine
    fmt.Println(divide(10, 0)) // Causes a panic but is recovered
    fmt.Println("This will print after recovery")
}
