package main

import "fmt"

func productExceptSelf(nums []int) []int {
	product := 1
	for _, num := range nums {
		product *= num
	}

	result := make([]int, len(nums))
	for i := range nums {
		result[i] = product / nums[i]
	}
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{10, 20}))
	fmt.Println(productExceptSelf([]int{3, 27, 4, 2}))
}
