package main

import "fmt"

func addDigits(num int) int {
    if num < 0 || num > 2147483647 {
        return -1
    }

    if num >= 0 && num <= 9 {
        return num
    }

    var sum int = num % 10 // 1 9

    num = num / 10 // 169,999,999
    sum = sum + (num % 10) // 2

    num = num / 10 // 16,999,999
    sum = sum + (num % 10) // 3

    num = num / 10 // 1,699,999
    sum = sum + (num % 10) // 4

    num = num / 10
    sum = sum + (num % 10) // 5

    num = num / 10
    sum = sum + (num % 10) // 6

    num = num / 10
    sum = sum + (num % 10) // 7

    num = num / 10
    sum = sum + (num % 10) // 8

    num = num / 10
    sum = sum + (num % 10) // 9

    num = num / 10
    sum = sum + (num % 10) // 10

    if sum > 9 {
        num = sum / 10
        sum = (sum % 10) + (num % 10)
    }

    if sum > 9 {
        num = sum / 10
        sum = (sum % 10) + (num % 10)
    }

    return sum
}


func main() {
    fmt.Println(addDigits(1699999999))
    fmt.Println(addDigits(67321))
    fmt.Println(addDigits(9))
}