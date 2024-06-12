package main

import "fmt"

func main() {
	var nama string
	nama = "Bayu"
	fmt.Println(nama)

	// atau

	var name = "Bayu"
	fmt.Println(name)

	//  atau
	asmane := "Hanisah"
	fmt.Println(asmane)

	// multi variable

	var (
		firstName = "Bayu"
		lastName  = "Hanisah"
	)

	fmt.Println(firstName + " " + lastName)
}
