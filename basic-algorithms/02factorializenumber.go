//https://play.golang.org/p/ffJ2BL_NrQ
package main

import (
	"fmt"
)

func factorialize(num int64) int64 {
	var bnoden, factor int64 = 1, 1

	for ; factor <= num; factor++ {
		bnoden *= factor
	}
	return bnoden
}

// Test
func main() {
	fmt.Println(factorialize(5))  // 120
	fmt.Println(factorialize(10)) // 3628800
	// This is wrong if you use int
	fmt.Println(factorialize(20)) // 2432902008176640000
	fmt.Println(factorialize(0))  // 1
}
