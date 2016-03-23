package main

import "fmt"

func main() {
    // This same code in PHP or Ruby returns 7
    // In Go it returns.. 8!
    fmt.Println(int((0.7 + 0.1) * 10))
}
