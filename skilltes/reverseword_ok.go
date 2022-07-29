package main

import (
	"fmt"
	"strings"
)

func main() {
	v := "I am A Great human"
	values := strings.Split(v, " ")
	for i, val := range values {
		values[i] = reverse(val)
	}
	fmt.Print(strings.Join(values, " "))
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)

}
