// https://play.golang.org/p/zDsYf6eTJ7

// Many possible ways for Go to handle JavaScript's mixed-type arrays, including map and json.Unmarshal.
/*
  I made all the arrays string arrays which IMHO is sufficient for the purposes of what was
  originally a basic JavaScript exercise. As we go deeper we'll probably have to change more
  solution requirements to address the differences between JS and Go.
*/

package main

import (
	"fmt"
)

func slasher(arr []string, howMany int) []string {
	if howMany > len(arr) {
		howMany = len(arr)
	}
	bnoden := arr[howMany:len(arr)]
	return bnoden
}

// Test
func main() {
	arr1 := []string{"1", "2", "3"}
	arr2 := []string{"burgers", "fries", "shake"}
	arr3 := []string{"1", "2", "chicken", "3", "potatoes", "cheese", "4"}

	fmt.Println(slasher(arr1, 2)) // [3]
	fmt.Println(slasher(arr1, 0)) // [1 2 3]
	fmt.Println(slasher(arr1, 9)) // []
	fmt.Println(slasher(arr1, 4)) // []
	fmt.Println(slasher(arr2, 1)) // [fries shake]
	fmt.Println(slasher(arr3, 5)) // [cheese 4]
}
