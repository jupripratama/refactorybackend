package main

import (
	"fmt"
)

func main() {
	var terms, next int
	var previous, current int = 0, 1
	fmt.Printf("\nMasukkan jumlah terms: ")
	fmt.Scanf("%d", &terms)
	if terms < 2 {
		fmt.Printf("\nJumlah minimal term pada Fibonacci series adalah dua")
	} else {
		fmt.Printf("\n\nFibonacci series:\n\n")
		for i := 0; i < (terms); i++ {
			fmt.Printf("%d ", previous)
			next = previous + current
			previous = current
			current = next
		}
	}
	fmt.Printf("\n")
}
