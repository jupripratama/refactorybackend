package main

import (
	"fmt"
	"strings"
)

func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "abc"
	str2 := "ABC"
	str3 := "Hello World"

	// Displaying strings
	fmt.Println("Strings (before):")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)
	fmt.Println("String 3:", str3)

	// Converting all the string into uppercase
	// Using ToUpper() function
	res1 := strings.ToUpper(str1)
	res2 := strings.ToUpper(str2)
	res3 := strings.ToUpper(str3)

	// Displaying the results
	fmt.Println("\nStrings (after):")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("Result 3:", res3)
}
