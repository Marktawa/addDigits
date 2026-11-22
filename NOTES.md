## Add the digits of an integer until they become a single digit

## Trial 1

```go
package main

import "fmt"

func addDigits(num int) int {
    var a int = num % 10
    var b int = num / 10

    if b == 0 {
        b = a
    }

    if b < 10 {
        return a
    }

    a = a + (b % 10)

    b = b / 10

    if b == 0 {
        b = a
    }

    if b < 10 {
        return a
    }

    a = a + (b % 10)
    b = b / 10

    if b == 0 {
        b = a
    }

    if b < 10 {
        return a
    }

    a = a + (b % 10) 
    b = b / 10

    if b == 0 {
        b = a
    }

    if b < 10 {
        return a
    }

    a = a + (b % 10) 
    b = b / 10
    
    if b == 0 && a < 10 {
        return a
    }
}
```

## Trial 2 (2 digit numbers)
```go
package main

import "fmt"

func addDigits(num int) int {
    var a int = num % 10
    var b int = num / 10

    a = a + (b % 10)

    // The highest possible single digit number is 9. The program continues running only if a is more than 9
    if (a > 9) {
        b = a / 10
        a = a % 10
        a = a + (b % 10)
    }

    return a
}

func main() {
    fmt.Println(addDigits(9))  // 9
    fmt.Println(addDigits(12)) // 3
    fmt.Println(addDigits(99)) // 9
}
```

<!-- Added bonus that this also works for single digit numbers as well. No need to implement code to handle single digit numbers-->

## Trial 3 (3 digit numbers)

```go
package main

import "fmt"

func addDigits(num int) int {
    var a int = num % 10
    var b int = num / 10

    a = a + (b % 10)

    b = b / 10

    a = a + (b % 10)

    if a > 9 {
        b = a / 10
        a = a % 10
        a = a + (b % 10)
    }

    if a > 9 {
        b = a / 10
        a = a % 10
        a = a + (b % 10)
    }

    return a
}

func main() {
    fmt.Println(addDigits(123))
    fmt.Println(addDigits(991))
}
```

## Trial 4 (5 digit numbers)

```go
package main

import "fmt"

func addDigits(num int) int {
    var a int = num % 10
    var b int = num / 10

    a = a + (b % 10)

    b = b / 10

    a = a + (b % 10)

    b = b / 10

    a = a + (b % 10)

    b = b / 10

    a = a + (b % 10)
    
    b = b / 10

    a = a + (b % 10)
    
    if a > 9 {
        b = a / 10
        a = a % 10
        a = a + (b % 10)
    }

    if a > 9 {
        b = a / 10
        a = a % 10
        a = a + (b % 10)
    }

    return a
}

func main() {
    fmt.Println(99993)
    fmt.Println(2037)
    fmt.Println(919)
    fmt.Println(49)
    fmt.Println(13)
    fmt.Println(6)
}
```

## Hypothesis (Brainstorm)

```
10 = 1 + 0 = 1
11 = 1 + 1 = 2
12 = 1 + 2 = 3
13 = 1 + 3 = 4
14 = 1 + 4 = 5
15 = 1 + 5 = 6
16 = 1 + 6 = 7
17 = 1 + 7 = 8
18 = 1 + 8 = 9
19 = 1 + 9 = 10 = 1 + 0 = 1
20 = 2 + 0 = 2
21 = 2 + 1 = 3
22 = 2 + 2 = 4
23 = 2 + 3 = 5
24 = 2 + 4 = 6
25 = 2 + 5 = 7
26 = 2 + 6 = 8
27 = 2 + 7 = 9
28 = 2 + 8 = 10 = 1 + 0 = 1
29 = 2 + 9 = 11 = 1 + 1 = 2
```

The sums are a repetition of the following sequence:
1,2,3,4,5,6,7,8,9. Between 10 and 99 the starts at 1 and ends at 9

## Hypothesis

```
100 = 1
101 = 2
102 = 3
103 = 4
104 = 5
105 = 6
106 = 7
107 = 8
108 = 9
109 = 10 = 1
110 = 2
```

## Max Digits (10 Digits)

```go
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
```

## Max Digits (Clean Up)

```go
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
```

The maximum possible number of digits that an integer has is 10. So we can extract all 10 digits while adding each digit to the variable that keeps track of the sum of digits. Regardless of whether the number has 10 digits or not we run the extraction code 10 times. Why? Because `0 / 10 = 0` and `0 % 10 = 0`. Our algorithm for keeping track of the sum adds the extracted number to the previous sum. Adding 0 to a number doesn't change the answer so we can confidently do it over and over again. No need to perform a check like `if b = 0`, as we definitely know that after 10 cycles of extracting the number, b = 0.

E.g. Cycles of dividing by 10 for a 10-digit number

```
1. 1,699,999,999 / 10 = 169,999,999
2. 169,999,999 / 10 = 16,999,999
3. 16,999,999 / 10 = 1,699,999
4. 1,699,999 / 10 = 169,999
5. 169,999 / 10 = 16,999
6. 16,999 / 10 = 1,699
7. 1,699 / 10 = 169
8. 169 / 10 = 16
9. 16 / 10 = 1
10. 1 / 10 = 0
```

The sum of the digits of an integer will not exceed a two-digit number if the integer is between 0 and 2^31 (0 < num < 2^31)

```
1,699,999,999 when its digits are summed gives 79

When 79's digits are summed it will give 16

16's digits will give 7

1,699,999,999 -> 79 -> 16 -> 7
```

A 2-digit number when its digits are summed can at most produce another 2-digit number once. The 2-digit number it produces if its digits are summed can only produce a single digit number. Hence you only need to check if the sum has exceeded 9 twice.

<!--
The number of cycles to add a 2-digit number will not exceed 2 for it to become a single digit. The highest possible sum for the digits of a 2-digit number is 18, which sums to a single digit, 9.
-->

```
99 -> 18 -> 9
```

If the number is divisible by 9, then the sum of the digits will give 9.
```go
if num % 9 == 0 {
    return 9
}
```

The highest possible sum for each type of number (classified according to number of digits). The highest possible sum of digits that will produce a 2-digit number whose sum of digits will also produce a 2-digit number.

```
10-digit num -> 1,699,999,999 -> 79 -> 16 -> 7
9-digit num -> 999,999,997 -> 79 -> 16 -> 7
8-digit num -> 99,999,996 -> 69 -> 15 -> 6
7-digit num -> 9,999,995 -> 59 -> 14 -> 5
6-digit num -> 999,994 -> 49 -> 13 -> 4
5-digit num -> 99,993 -> 39 -> 12 -> 3
4-digit num -> 9,992 -> 29 -> 11 -> 2
3-digit num -> 991 -> 19 -> 10 -> 1
2-digit num -> 99 -> 18 -> 9
1-digit num -> 9
```

### Hypothesis: Beyond the Limit

```
9,999,999,998
99,999,999,999
999,999,999,990 -> 12 digits -> 99
9,999,999,999,900 -> 13 digits -> 99

999,999,999,999,999,999,999 -> 21 digits -> 189
9,999,999,999,999,999,999,999 -> 22 digits -> 198
99,999,999,999,999,999,999,991 -> 23 digits -> 199
999,999,999,999,999,999,999,910 -> 24 digits -> 199
```

Sum of digits beyond
```
199 -> 19 -> 10 -> 1 -> 9 x 22
289, 298 -> 19 -> 10 -> 1 -> 9 x 32
379, 388, 398 -> 19 -> 10 -> 1 -> 9 x 42
469, 478, 487, 496 -> 19 -> 10 -> 1 -> 9 x 52
```

Lowest possible
```
199 -> 19 -> 10 -> 1
919 -> 19 -> 10 -> 1
991 -> 19 -> 10 -> 1

289 -> 19 -> 10 -> 1

199, 919, 991
289, 298, 829, 892, 928, 982
379, 397, 937, 973, 739, 793
469, 496, 946, 964, 649, 694
559, 595, 955
```

The highest number to produce a 2-digit number that produces a 2-digit number, isn't the highest possible number. It's a number that is a combination of all the digits being 9 except for 1. That digit, say `x = n - 2`, where `x` represents the value of the digit and `n` is the number of digits in the number of interest.

`10^n - (n - 1) or 10^n - n + 1`

```
10^n - 10 + (n - 2)
```

```
n = 9, 79 = 72 + 7 = 8 x 9 + 7 = 9(n - 1) + (n - 2)
```