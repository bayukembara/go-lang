package main

import "fmt"

func main() {
	fmt.Println("this is function")
	sayHello()
	sayHelloTo("Bayu", "Hanisah")

	result := getHello("Hanisah")
	fmt.Println(result)

	firstName, lastName := getACouple()
	fmt.Println(firstName, lastName)

	//! ketika ingin mengambil salah satu saja

	firstThink, _ := getACouple()
	fmt.Println(firstThink)

	lastLongCouple := getCoupleMarried("Bayu", "Hanisah")
	fmt.Println(lastLongCouple)

	//! func arg
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	//! change function from slice
	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

	goodbye := getGoodBye("Vintage")
	fmt.Println(goodbye)
}

func sayHello() {
	fmt.Println("Hello from function sayHello()")
}

//! func Parameter

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName+" "+lastName)
}

//! func return parameter

func getHello(name string) string {
	return "Hello" + " " + name
}

//! func returning multiple values

func getACouple() (string, string) {
	return "Bayu", "Hanisah"
}

//! func returning named values

func getCoupleMarried(husbandName, wifeName string) string {
	return "Here is the married couple: husband's name is " + husbandName + " and the wife's name is " + wifeName
}

//! function slice values
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//! function becomes variable and values

func getGoodBye(name string) string {
	return "Good Bye " + name
}
