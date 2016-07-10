// https://play.golang.org/p/DmMymFUaNw
package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func palindrome(str string) bool {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatalln()
	}

	cleanstr := strings.ToLower(reg.ReplaceAllString(str, ""))
	arr := []rune(cleanstr)
	bnoden := []rune{}
	isPalindrome := false
	count := 0

	for i := len(arr) - 1; i >= 0; i-- {
		bnoden = append(bnoden, arr[i])
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == bnoden[i] {
			count++
		}
	}
	if count == len(arr) {
		isPalindrome = true
	}
	return isPalindrome
}

// Test
func main() {
	fmt.Println(palindrome("eye"))                            // true
	fmt.Println(palindrome("_eye"))                           // true
	fmt.Println(palindrome("race car"))                       // true
	fmt.Println(palindrome("not a palindrome"))               // false
	fmt.Println(palindrome("A man, a plan, a canal. Panama")) // true
	fmt.Println(palindrome("never odd or even"))              // true
	fmt.Println(palindrome("nope"))                           // false
	fmt.Println(palindrome("almostomla"))                     // false
	fmt.Println(palindrome("My age is 0, 0 si ega ym."))      // true
	fmt.Println(palindrome("1 eye for of 1 eye."))            // false
	// I used backticks instead of quotes to handle "unknown escape sequences". Still figuring out best way to deal with it.
	fmt.Println(palindrome(`0_0 (: /-\ :) 0-0`)) // true
	fmt.Println(palindrome(`five|\_/|four`))     // false
}
