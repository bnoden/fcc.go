// https://play.golang.org/p/ouLXhwuMbV
package main

import (
	"fmt"
)

func factorialize(num int64) int64 {
	var bnoden int64 = 1
	var i int64

	for i = 1; i <= num; i++ {
		bnoden *= i
	}
	return int64(bnoden)
}

// Test
func main() {
	fmt.Println(factorialize(5))
	fmt.Println(factorialize(10))
	fmt.Println(factorialize(20))
	fmt.Println(factorialize(0))
}