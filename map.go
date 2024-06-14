package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John",
		"address": "British",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// function map

	fmt.Println(len(person))

	book := map[string]string{
		"title":  "Golang Books",
		"author": "Bayu Kembara",
		"wrong":  "Upss",
	}

	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}
