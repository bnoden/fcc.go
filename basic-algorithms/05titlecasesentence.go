// https://play.golang.org/p/4hDLNm25Kv
package main

import (
    "fmt"
    "strings"
)

func titleCase(str string) string {
    arr := strings.Fields(str)
    var bnoden []string = arr[0:len(arr)] // Slices are the key to working with arrays in Go.
    
    for i := 0; i < len(arr); i++ {
        bnoden[i] = strings.ToUpper(arr[i][:1]) + strings.ToLower(arr[i][1:])
    }
    return strings.Join(bnoden, " ")
}
// Test
func main() {
    fmt.Println(titleCase("I'm a little tea pot")) // I'm A Little Tea Pot
    fmt.Println(titleCase("sHoRt AnD sToUt")) // Short And Stout
    fmt.Println(titleCase("HERE IS MY HANDLE HERE IS MY SPOUT")) // Here Is My Handle Here Is My Spout
}