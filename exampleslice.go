package main

import "fmt"

func main() {
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	daySlice1 := days[5:]
	daySlice1[0] = "new saturday"
	daySlice1[1] = "new sunday"

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("\n-----------------------------APPEND SLICE------------------------")

	// Make Slice

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Bayu"
	newSlice[1] = "Bayu"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	// how to add new array in append

	newSlice2 := append(newSlice, "Bayu")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Bay"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("\n----------------------------COPY SLICE-----------------------")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fmt.Println("\n----------------------------DIFFERENT ARRAY AND SLICE-----------------------")
	iniArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	iniSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
