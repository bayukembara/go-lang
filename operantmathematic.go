package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b

	fmt.Println(c)

	// augmented operation

	var i = 10

	i += 10 // i = i(lama(10)) + 10
	fmt.Println(i)

	i += 5 // i = i(lama(20)) + 5
	fmt.Println(i)

	// unary operand
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
