package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Daryn"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting message and print it.
	// message := greetings.Hello("Gladys")
	fmt.Println(messages)
}
