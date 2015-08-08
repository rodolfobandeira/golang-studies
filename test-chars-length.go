package main

import "fmt"
import "unicode/utf8"

func main() {
    fmt.Println("Testando caractere: ç", len("ç"), utf8.RuneCountInString("ç"))
}

// Testando caractere: ç 2 1
