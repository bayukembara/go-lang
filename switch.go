package main

import "fmt"

func main() {
	name := "Hanisah"

	switch name {
	case "Bayu":
		fmt.Println(name)
	case "Hanisah":
		fmt.Println(name)
	default:
		fmt.Println("Boleh Beritahukan nama anda?")
	}

	// Switch using condition
	switch length := len(name); length < 5 {
	case true:
		fmt.Println("Perlu lebih baik lagi")
	case false:
		fmt.Println("Hanisah terbaik untuk diriku")
	}

	// Switch without condition

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Jangan sakitin hati hanisah")
	case length == 5:
		fmt.Println("Hanisah sudah sangat komplit untuk kamu")
	case length < 5:
		fmt.Println("Jangan cemas pasti hanisah amanah")
	}
}
