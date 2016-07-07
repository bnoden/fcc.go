// https://play.golang.org/p/meKL28q9wY
package main

import (
	"fmt"
)

func confirmEnding(str, target string) bool {
	bnoden := false

	if str[len(str)-len(target):len(str)] == target {
		bnoden = true
	}
	return bnoden
}

// Test
func main() {
	fmt.Println(confirmEnding("Bastian", "n"))                                                                                                         // true
	fmt.Println(confirmEnding("Connor", "n"))                                                                                                          // false
	fmt.Println(confirmEnding("Walking on water and developing software from a specification are easy if both are frozen", "specification"))           // false
	fmt.Println(confirmEnding("He has to give me a new name", "name"))                                                                                 // true
	fmt.Println(confirmEnding("He has to give me a new name", "me"))                                                                                   // true
	fmt.Println(confirmEnding("He has to give me a new name", "na"))                                                                                   // false
	fmt.Println(confirmEnding("If you want to save our world, you must hurry. We dont know how much longer we can withstand the nothing", "mountain")) // false
}
