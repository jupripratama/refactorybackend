package main

import "fmt"

func printSequence(x int) int {

	t1 := 0
	t2 := 1
	temp := 0

	for i := 1; i <= x; i++ {

		if i == 1 {
			fmt.Print(t1, ", ")
			continue
		}
		if i == 2 {
			fmt.Print(t2, ", ")
			continue
		}
		temp = t1 + t2
		t1 = t2
		t2 = temp

		fmt.Print(temp)
		if i != x {
			fmt.Print(", ")
		}

	}
	fmt.Println()

	return temp
}

func main() {

	var x int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&x)

	max := printSequence(x)
	fmt.Println("the maximum input is: ", max)

}
