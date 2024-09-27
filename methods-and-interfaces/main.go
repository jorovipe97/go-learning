package main

import "fmt"

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
// In this example, both Scale and Abs are methods with receiver type *Vertex, even though the Abs method needn't modify its receiver.
// https://go.dev/tour/methods/8
func (v Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// You can only declare a method with a receiver whose type is
// defined in the same package as the method. You cannot declare
// a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int).

type MyFloat float64 // We define a type in the same package.

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	v := Vertex{1, 1}
	fmt.Println(v.Abs())
	fmt.Println(v)
	v.Scale(5)
	fmt.Println(v)

	var abser Abser = v
	describe(abser)

	// fmt.Println(MyFloat(-5).Abs())
	// fmt.Println(MyFloat(-2).Abs())
	// fmt.Println(MyFloat(3).Abs())
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
