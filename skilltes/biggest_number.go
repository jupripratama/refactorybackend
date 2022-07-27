package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func maxRedigit(num int64) int64 {
	numArr := strings.Split(strconv.FormatInt(num, 10), "")
	sort.Slice(numArr, func(i, j int) bool {
		return numArr[i] > numArr[j]
	})
	res, _ := strconv.ParseInt(strings.Join(numArr, ""), 10, 64)
	return res
}
func main() {
	var num int64 = 1000
	fmt.Print("Masukan Number : ")
	fmt.Scan(&num)
	if num < 100 || num > 999 {
		fmt.Print("nill")
	} else {
		println(maxRedigit(num))
	}

}
