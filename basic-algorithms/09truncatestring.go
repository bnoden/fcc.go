// https://play.golang.org/p/EzvhWMljku
package main

import (
	"fmt"
)

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}

// Test
func main() {
	str1 := "A-tisket a-tasket A green and yellow basket"
	str2 := "Peter Piper picked a peck of pickled peppers"

	fmt.Println(truncateString(str1, 11))               // "A-tisket..."
	fmt.Println(truncateString(str2, 14))               // "Peter Piper..."
	fmt.Println(truncateString(str1, len(str1)))        // "A-tisket a-tasket A green and yellow basket"
	fmt.Println(truncateString(str1, len(str1)+2))      // "A-tisket a-tasket A green and yellow basket"
	fmt.Println(truncateString("A-", 1))                // "A..."
	fmt.Println(truncateString("Absolutely Longer", 2)) // "Ab..."
}
