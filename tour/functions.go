package main

import "fmt"

// type comes after name
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}

