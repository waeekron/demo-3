package main

import "fmt"

func main() {
	fmt.Println("Enter a string with spaces:")

	var input string
	fmt.Scanln(&input)

	fmt.Println("You entered:", input)
}
