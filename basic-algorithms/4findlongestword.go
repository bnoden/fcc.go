// https://play.golang.org/p/hXf_wdZdrC
package main

import (
	"fmt"
	"strings"
)

func findLongestWord(str string) int {
	bnoden := 1
	arr := strings.Fields(str)

	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > bnoden {
			bnoden = len(arr[i])
		}
	}
	return bnoden
}

// Test
func main() {
	fmt.Println(findLongestWord("The quick brown fox jumped over the lazy dog"))                 // 6
	fmt.Println(findLongestWord("May the force be with you"))                                    // 5
	fmt.Println(findLongestWord("Google do a barrel roll"))                                      // 6
	fmt.Println(findLongestWord("What is the average airspeed velocity of an unladen swallow"))  // 8
	fmt.Println(findLongestWord("What if we try a super-long word such as otorhinolaryngology")) // 19
}
