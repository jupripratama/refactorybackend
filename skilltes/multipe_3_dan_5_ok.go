package main

import "fmt"

func solution(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(solution(10))
	fmt.Println(solution(20))
}
