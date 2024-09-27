package main

import "fmt"

type Position struct {
	X int
	Y int
}

func structs() {
	pos := Position{1, 1}
	posPointer := &pos
	fmt.Println(pos, posPointer)
	modifyStruct(posPointer)
	fmt.Println(pos, posPointer)
	passStruct(pos)
	fmt.Println(pos, posPointer)
}

func modifyStruct(pos *Position) {
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	pos.X = 10
	(*pos).Y = 10 // Equivalent but cumbersome.
}

func passStruct(pos Position) {
	// Argument is passed as copy. Modifications here do not reflect in the caller.
	pos.X = 20
	pos.Y = 20
}