package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // Hides the time, source file, and line number.

	// A slice of names.
	names := []string{"Luisa", "Jose", "Miel"}

	// Get a greeting message and print it.
	message, error := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if error != nil {
		log.Fatal(error)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
