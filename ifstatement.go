package main

import "fmt"

func main() {
	name := "Budi"

	if name == "Bayu" {
		fmt.Println("Hello" + " " + name)
	} else {
		fmt.Println("What's your name?")
	}
}
