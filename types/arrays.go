package main

import (
	"fmt"
	"strings"
)

func arrays() {
	// The type [n]T is an array of n values of type T.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	var s2 []int = primes[2:4]
	fmt.Println(s, s2)

	s[1] = 100
	fmt.Println(s, s2)
	fmt.Println(len(s), cap(s))
	fmt.Println(len(s2), cap(s2))

	// Extend over the length of the slice
	// Out of bounds exception.
	// s3 := primes[:7]
	// fmt.Println(s3)

	// Creates an array an then a slice from it.
	s4 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s4, len(s4), cap(s4))

	// https://go.dev/tour/moretypes/13
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	s5 := make([]int, 2, 10) // 2 length, 10 capacity
	fmt.Println(s5, len(s5), cap(s5))
	s5 = append(s5, 1)
	fmt.Println(s5, len(s5), cap(s5))
	s5_2 := []int{ 2, 3, 4 }
	s5 = append(s5, s5_2...)
	fmt.Println(s5, len(s5), cap(s5))

	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		{"_", "_", "_"}, // Can also be written like this
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
