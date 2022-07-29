package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var (
		n             = len(nums)
		leftProducts  = make([]int, n)
		rightProducts = make([]int, n)
	)
	leftProducts[0] = 1
	rightProducts[len(nums)-1] = 1

	for i := 0; i < n-1; i++ {
		leftProducts[i+1] = leftProducts[i] * nums[i]
		rightProducts[n-i-2] = rightProducts[n-i-1] * nums[n-i-1]
	}

	result := make([]int, n)
	for i := range result {
		result[i] = leftProducts[i] * rightProducts[i]
	}
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{12, 20}))
	fmt.Println(productExceptSelf([]int{3, 27, 4, 2}))
	fmt.Println(productExceptSelf([]int{13, 10, 5, 2, 9}))
	fmt.Println(productExceptSelf([]int{16, 17, 4, 3, 5, 2}))
}
