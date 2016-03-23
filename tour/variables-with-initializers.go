package main

import "fmt"

var i, j int = 1, 2
// i = 1
// j = 2

func main() {
    var c, python, java = true, false, "no!"
    // c      = true and automatically signed to boolean
    // python = false and automatically signed to boolean
    // java   = "no" and automatically signed to string

    fmt.Println(i, j, c, python, java)
    // 1 2 true false no
}

