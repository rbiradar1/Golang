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
	log.SetPrefix("Welcome, hi: ")
	log.SetFlags(0)

	//message := greetings.Hello("Ruchita")
	message, err := greetings.Hello("") //error handling in case of empty string
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
