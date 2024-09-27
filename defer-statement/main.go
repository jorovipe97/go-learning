package main

import (
	"fmt"
)

// https://go.dev/blog/defer-panic-and-recover
// func main() {
// 	defer fmt.Println("This was deffered 1.")
// 	defer fmt.Println("This was deffered 2.")
// 	fmt.Println("Hello, World!")
// 	fmt.Println("Hello, World 2!")
// 	defer fmt.Println("This was deffered 3.")
// 	fmt.Println("Hello, World 3!")
// }

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}