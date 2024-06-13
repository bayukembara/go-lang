package main

import "fmt"

func main() {
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	daySlice1:= days[5:]
	daySlice1[0] = "new saturday"
	daySlice1[1] = "new sunday"

	fmt.Println(days)
}
