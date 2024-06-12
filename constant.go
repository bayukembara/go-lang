package main

import "fmt"

func main() {
	const firstName = "Bayu"
	const lastName = "Hanisah"

	fmt.Println(firstName + " " + lastName)

	// Multiple

	const (
		fName = "Bayu"
		lName = "Hanisah"
	)

	fmt.Println(fName + " " + lName)
}
