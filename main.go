package main

import "fmt"

func addDigits(num int) int {
    var digit1 int = num % 10
    var temp int = num / 10

    var digit2 int = temp % 10
    var digit3 int = temp / 10

    var sum int = digit1 + digit2 + digit3

    if sum > 9 {
        digit1 = sum % 10
        digit2 = sum / 10
        sum = digit1 + digit2
    }

    if sum > 9 {
        digit1 = sum % 10
        digit2 = sum / 10
        sum = digit1 + digit2
    }

    return sum
}

func main() {
    fmt.Println(addDigits(289))
    fmt.Println(addDigits(38))
    fmt.Println(addDigits(0))
}