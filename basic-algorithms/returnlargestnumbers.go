// https://play.golang.org/p/Mf8jkcp0gR
package main

import (
	"fmt"
)

func largestOfFour(arr [4][4]int) [4]int {
	bnoden := [4]int{0, 0, 0, 0}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > bnoden[i] {
				bnoden[i] = arr[i][j]
			}
		}
	}
	return bnoden
}

// Test
func main() {
	arr1 := [4][4]int{{13, 27, 18, 26}, {4, 5, 1, 3}, {32, 35, 37, 39}, {1000, 1001, 857, 1}}
	arr2 := [4][4]int{{4, 9, 1, 3}, {13, 35, 18, 26}, {32, 35, 97, 39}, {1000000, 1001, 857, 1}}

	fmt.Println(largestOfFour(arr1)) // [27 5 39 1001]
	fmt.Println(largestOfFour(arr2)) // [9 35 97 1000000]
}
