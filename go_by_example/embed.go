package main

import (
   _ "embed"
    "fmt"
)

//go:embed hello.txt
var content string

func main() {
    fmt.Println(content)
}
