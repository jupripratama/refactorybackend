package main

import "fmt"

func extractMiddle(str string) string {

	var position, length int

	if len(str)%2 == 1 {
		position = len(str) / 2
		length = 1
	} else {
		position = len(str)/2 - 1
		length = 2
	}

	return str[position : position+length]
}

func main() {
	fmt.Println(extractMiddle("QRSTU"))
	fmt.Println(extractMiddle("RSTU"))
	fmt.Println(extractMiddle("TUVWXYZ"))
	fmt.Println(extractMiddle("QRSTUVWXYZ"))
}
