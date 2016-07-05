package main
// Needs work






import (
    "fmt"
    "regexp"
    "strings"
)

func palindrome(str string) bool {
    var bnoden string = strings.NewReplacer(`w, "")
    result := bnoden.Replace(str)
    return result
}
// Test
func main() {
    fmt.Println(palindrome("eye"))
    fmt.Println(palindrome("_eye"))
    fmt.Println(palindrome("race car"))
    fmt.Println(palindrome("not a palindrome"))
    fmt.Println(palindrome("A man, a plan, a canal. Panama"))
    fmt.Println(palindrome("never odd or even"))
    fmt.Println(palindrome("nope"))
    fmt.Println(palindrome("almostomla"))
    fmt.Println(palindrome("My age is 0, 0 si ega ym."))
    fmt.Println(palindrome("1 eye for of 1 eye."))
    fmt.Println(palindrome("0_0 (: /-\ :) 0-0"))
    fmt.Println(palindrome("five|\_/|four"))
}