// https://play.golang.org/p/H1_W6gcdgZ

package main

import (
	"fmt"
)

func reverseString(str string) string {
	arr := []rune(str)
	bnoden := []rune{}

	for i := len(arr) - 1; i >= 0; i-- {
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
