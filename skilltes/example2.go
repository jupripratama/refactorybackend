package main

import "fmt"

func main() {
	var Abjad = []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"x",
		"Y",
		"Z"}
	fmt.Println(Abjad[16:20:22]) // [apple, grape]
	fmt.Println(Abjad[17:21])    // [apple, grape]

}
