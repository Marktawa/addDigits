package main

import "fmt"

func addDigits(num int) int {
    if num >= 0 && num <= 9 {
        return num
    }
    return -1
}

func main() {
    fmt.Println(addDigits(9))
    fmt.Println(addDigits(3))
    fmt.Println(addDigits(0))
    fmt.Println(addDigits(11))
}