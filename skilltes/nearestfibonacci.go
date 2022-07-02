package main

import "fmt"

func main() {
    s := "test123 4"
    fmt.Println(reverseString(s))
}

func reverseString(s string) string {
    a := []byte(s)
    for i, j := 0, len(s)-1; i < j; i++ {
        a[i], a[j] = a[j], a[i]
        j--
    }
    return string(a)
}
