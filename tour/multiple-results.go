package main

import "fmt"

// Function returning multiple results
// Thats cool!

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("rodolfo", "bandeira")
    fmt.Println(a, b)
}

