// https://play.golang.org/p/3XWkf0h_Xw
package main

import (
	"fmt"
	"strings"
)

func repeatStringNumTimes(str string, num int) string {
	bnoden := ""
	if num > 0 {
		bnoden = strings.Repeat(str, num)
	}
	return bnoden
}

// Test
func main() {
	fmt.Println(repeatStringNumTimes("*", 3))    // ***
	fmt.Println(repeatStringNumTimes("abc", 3))  // abcabcabc
	fmt.Println(repeatStringNumTimes("abc", 4))  // abcabcabcabc
	fmt.Println(repeatStringNumTimes("abc", 1))  // abc
	fmt.Println(repeatStringNumTimes("*", 8))    // ********
	fmt.Println(repeatStringNumTimes("abc", -2)) // ""
}
