package main

import "fmt"

func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
        fmt.Println(sum)
    }
    fmt.Println(sum)
}

/*
2
4
8
16
32
64
128
256
512
1024
1024
*/

