package main

import "fmt"

func main() {
	type noKTP string

	var ktpB noKTP = "111111111"
	fmt.Println(ktpB)
	fmt.Println(noKTP("2222222222"))

	// conversion

	var contoh string = "22222222222"
	var contohKtp noKTP = noKTP(contoh)

	fmt.Println(ktpB)
	fmt.Println(contohKtp)
}
