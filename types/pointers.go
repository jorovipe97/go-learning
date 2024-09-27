package main

import "fmt"

func pointers()  {
	i := 42
	// The type *T is a pointer to a T value. Its zero value is nil.
	// The & operator generates a pointer to its operand.
	var a *int = &i
	fmt.Println(a, *a, i)

	*a = 50
	fmt.Println(a, *a, i)
}
