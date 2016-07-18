// https://play.golang.org/p/apC1p7ubf3
package main

import (
	"fmt"
)

func reverseString(str string) string {
	arr, bnoden := []rune(str), []rune{}

	for i := len(arr)-1; i >= 0; i-- {
		bnoden = append(bnoden, arr[i])
	}
	return string(bnoden)
}

// Test
func main() {
	fmt.Println(reverseString("hello"))                // "olleh"
	fmt.Println(reverseString("Howdy"))                // "ydwoH"
	fmt.Println(reverseString("Greetings from Earth")) // "htraE morf sgniteerG"

}
