package main

import "fmt"

func addDigits(num int) int {
    var a int = num % 10 // 1st Digit

    var b int = num / 10
    var c int = b % 10
    a = a + c // 2nd

    b = b / 10
    c = b % 10
    a = a + c // 3rd

    b = b / 10
    c = b % 10
    a = a + c // 4th

    b = b / 10
    c = b % 10
    a = a + c // 5th

    b = b / 10
    c = b % 10
    a = a + c // 6th

    b = b / 10
    c = b % 10
    a = a + c // 7th

    b = b / 10
    c = b % 10
    a = a + c // 8th

    b = b / 10
    c = b % 10
    a = a + c // 9th

    b = b / 10
    c = b % 10
    a = a + c // 10th

    if a > 9 {
        b = a / 10
        a = a % 10
        c = b % 10
        a = a + c
    }

    if a > 9 {
        b = a / 10
        a = a % 10
        c = b % 10
        a = a + c
    }

    return a
}

func main() {
    fmt.Println(addDigits(1699999999))
    fmt.Println(addDigits(67321))
    fmt.Println(addDigits(9))
}