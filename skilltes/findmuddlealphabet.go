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
	// fmt.Println(Abjad[16:21])
	// fmt.Println("Found:", Abjad[17:21])
	// fmt.Println("Q and Z:", Abjad[19:20:26])
	// fmt.Println(Abjad[16:26])
	fmt.Println("Q and U:", Abjad[18])
	fmt.Println("R and U:", Abjad[18:20])
	fmt.Println("T and Z:", Abjad[22])
	fmt.Println("T and Z:", Abjad[20:22])

}
