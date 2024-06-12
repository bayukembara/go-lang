package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Bayu"
	names[1] = "Hanisah"

	fmt.Println(names)

	// or

	var values = [3]int{
		90, 80, 70,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[0])

	// array without limit must be declare if not it can't be used

	var arr = [...]int{
		90,
		90,
		90,
		90,
		90,
	}

	fmt.Println(arr)
	fmt.Println(len(arr))
}
