package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacciBuilderLong() func() int {
	prev1 := 1 // n - 1
	prev2 := 0 // n - 2
	return func () int {	
		current := prev2

		t2 := prev1
		t1 := prev1 + prev2
		prev2 = t2
		prev1 = t1
		return current
	}
}

// https://stackoverflow.com/questions/25491370/fibonacci-closure-in-go
// https://stackoverflow.com/questions/48836122/golang-multiple-assignment-evaluation
func fibonacciBuilderShort() func() int {
	prev2, prev1 := 0, 1
	return func () int {	
		current := prev2

		// First is evaluated from left-to-right the expressions on the left side.
		// Seconds is assigned to the variables from left-to-right on the right side.
		prev2, prev1 = prev1, prev1 + prev2
		return current
	}
}

func runFibonacci() {
	fibonacciLong := fibonacciBuilderShort()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciLong())
	}
}
