package main

import "fmt"

func main() {
	const (
		firstName string	= "Rifki"
		lastName 			= "Malaika"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Nurmanto"
	// lastName = "Malika"
}