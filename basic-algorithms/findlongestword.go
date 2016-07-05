// needs work

package main

import (
	"fmt"
	"strings"
)

func findLongestWord(str string) int {
	bnoden := 1
	arr := strings.Fields(str)

	for i := 0; i <= len(arr); i++ {
		if len(arr[i]) > bnoden {
			bnoden = len(arr[i])
		}
	}
	return bnoden
}

func main() {
	fmt.Println(findLongestWord("The quick brown fox jumped over the lazy dog"))
}
