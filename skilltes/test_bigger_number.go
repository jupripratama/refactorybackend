package main

import (
	"fmt"
	"sort"
	"strconv"
)

func maxRedigit(num int) int {

	if num < 100 || num > 999 {
		return -1

	}

	numStr := strconv.Itoa(num)
	numArr := []int{}

	for _, i := range numStr {
		num, _ := strconv.Atoi(string(i))
		numArr = append(numArr, num)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numArr)))

	newNum, _ := strconv.Atoi(fmt.Sprintf("%d%d%d", numArr[0], numArr[1], numArr[2]))

	return newNum
}

func main() {
	fmt.Println(maxRedigit(123))
	fmt.Println(maxRedigit(231))
	fmt.Println(maxRedigit(321))
	fmt.Println(maxRedigit(-1))
	fmt.Println(maxRedigit(0))
	fmt.Println(maxRedigit(99))
	fmt.Println(maxRedigit(1000))
}
