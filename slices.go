package main

import "fmt"

func main() {
	names := [...]string{"Bayu", "Hanisah", "Kembara", "Damayanti"}

	slice1 := names[2:4]
	slice2 := names[:4]
	slice3 := names[1:]
	slice4 := names[:]

	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice2))
	fmt.Println(append(slice1,slice2...))
}
