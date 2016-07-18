package main

import (
    "fmt"
    "strings"
)

func mutation(arr []string) []string {
    bnoden := false
    first := strings.ToLower(arr[0])
    next := strings.ToLower(arr[1])
    arrNext := strings.Fields(next)
    

    for i := 0; i < len(arrNext); i++ {
        if strings.Contains(first, arrNext[i]) {
            bnoden = true
        }
    }

    
    //fmt.Println(count)
    return arrNext
}

func main() {
    fmt.Println(mutation([]string{"hello", "hey"})) // false
    fmt.Println(mutation([]string{"hello", "Hello"})) // true
    fmt.Println(mutation([]string{"zyxwvutsrqponmlkjihgfedcba", "qrstu"})) // true
    fmt.Println(mutation([]string{"Mary", "Army"})) // true
    fmt.Println(mutation([]string{"Mary", "Aarmy"})) // true
    fmt.Println(mutation([]string{"Alien", "line"})) // true
    fmt.Println(mutation([]string{"floor", "for"})) // true
    fmt.Println(mutation([]string{"hello", "neo"})) // false
    fmt.Println(mutation([]string{"voodoo", "no"})) // false
}