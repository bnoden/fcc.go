// https://play.golang.org/p/HrFh0tQC-M
package main

import (
	"fmt"
)

func chunkArrayInGroups(arr []rune, size int) (bnoden [][]rune) {

	temp := []rune{}

	for i := 0; i < len(arr); i++ {
		temp = append(temp, arr[i])
		if i%size == size-1 {
			bnoden = append(bnoden, temp[0:len(temp)])
			temp = []rune{}
		}
	}
	if len(temp) != 0 {
		bnoden = append(bnoden, temp[0:len(temp)])
	}
	return bnoden
}

// Test
func main() {
	arr1 := []rune{'a', 'b', 'c', 'd'}
	arr2 := []rune{0, 1, 2, 3, 4, 5}
	arr3 := []rune{0, 1, 2, 3, 4, 5}
	arr4 := []rune{0, 1, 2, 3, 4, 5}
	arr5 := []rune{0, 1, 2, 3, 4, 5, 6}
	arr6 := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arr7 := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("%c\n", chunkArrayInGroups(arr1, 2)) // [[a b] [c d]]
	fmt.Println(chunkArrayInGroups(arr2, 3))        // [[0 1 2] [3 4 5]]
	fmt.Println(chunkArrayInGroups(arr3, 2))        // [[0 1] [2 3] [4 5]]
	fmt.Println(chunkArrayInGroups(arr4, 4))        // [[0 1 2 3] [4 5]]
	fmt.Println(chunkArrayInGroups(arr5, 3))        // [[0 1 2] [3 4 5] [6]]
	fmt.Println(chunkArrayInGroups(arr6, 4))        // [[0 1 2 3] [4 5 6 7] [8]]
	fmt.Println(chunkArrayInGroups(arr7, 2))        // [[0 1] [2 3] [4 5] [6 7] [8]]
}
