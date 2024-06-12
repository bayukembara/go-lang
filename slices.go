package main

import "fmt"

func main() {
	names := [...]string{"Bayu", "Hanisah", "Kembara", "Damayanti"}

	slice1 := names[2:4]
	slice2 := names[:4]

	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
