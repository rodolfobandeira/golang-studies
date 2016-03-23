package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x

    // How cool is that? passing just return will return both x and y
    return
}

func main() {
    fmt.Println(split(17))
}

