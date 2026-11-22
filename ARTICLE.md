# Coding Challenge: Add Digits using Go

This article goes over the solution to the following problem:

*Given an integer `num`, repeatedly add all its digits until the result has only one digit, and return it*

Source: LeetCode

**Example 1:**

```
Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
```

**Example 2:**
```
Input: num = 0
Output: 0
```

**Constraints:**
- `0 <= num <= 2^31 - 1`

**Follow up:** Could you do it without any loop/recursion in `O(1)` runtime?

## Prerequisites

For this solution I will use Go programming language (Golang). To follow along, you need to have it installed. 

## First Steps

Let's start with the easier stuff. We know that single digit numbers: 0, 1, 2, 3, 4, 5, 6, 7, 8, and 9 should just be returned as is.

The final solution should use a function called `addDigits`. We can set conditions that if a number is between 0 and 9 inclusive then the function should just return the same number as it is already a single digit. For any other number we will set it to return `-1` for now.

Create a file named `main.go` and add the following code:
```go
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
```

Save the file and run it in your terminal:
```sh
go run main.go
```

You should see the following output:
```
9
3
0
-1
```

## Two Digit Numbers

Now that we have handled single digit numbers, we can move on.

If a number has more than one digit, we need to first find a way to extract the digits of an integer. One way to do this is by making use of the `%` (pronounced modulo) operator. This gives the remainder when you divide a number by another number. For example, `5 / 2 = 2 remainder 1` therefore `5 % 2 = 1`.

If you find the modulo when a number is divided by 10, the answer you get is always the last digit of that number. <!-- Look for supporting theorems --> For example `35 % 10 = 5`. In our example `5` is the last digit of the number.

Before we proceed, let's test this out in Go. 

Create a file named `main.go` and add the following code:
```go
package main

import "fmt"

func main() {
    fmt.Println(35 % 10)
}
```

Save the file and run the code in your shell:
```shell
go run main.go
```

You should see `5` printed in your shell. 

<!-- The final solution should use a function called `addDigits`. So here's how it would look  like using that: -->

```go
package main

import "fmt"

func addDigits(num int) int {
    return num % 10
}

func main() {
    fmt.Println(addDigits(35))
}
```

Now that we have found a method to extract the last digit of a number accurately, we have a few more questions to address. How do we extract all its digits not just that the last digit? How do we add all the digits?

To address these questions, we need to simplify the approach to the problem. For now, we will only focus on 2-digit numbers (10 <= num <= 100). This makes it easier for us, as we only need to extract two digits then sum them. We already know how to extract the last digit, how then do we extract the first digit? 

To extract the next digit, just divide the number by 10.

For our example number 35, `35 / 10 = 3`

In integer division there are no decimal points and Go programming language floors the decimal value, hence we can confidently get the first digit simply by dividing the number by 10

<!-- First, we need to remove the digit we just extracted.

For example for the number 35, if we have identified 5 as the last digit, we need to remove from 35 so that it becomes 3. This will help us identify 3 as the next digit -->