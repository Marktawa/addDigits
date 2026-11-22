# Coding Challenge: Add Digits using Go

This article goes over the solution to the following problem:

*Given an integer `num`, repeatedly add all its digits until the result has only one digit, and return it*

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

To solve this problem, we need to first find a way to extract the digits of an integer. One way to do this is by making use of the `%` (pronounced modulo) operator. This gives the remainder when you divide a number by another number. For example, `5 / 2 = 2 remainder 1` therefore `5 % 2 = 1`.

If you find the modulo when a number is divided by 10, the answer you get is always the last digit of that number. <!-- Look for supporting theorems --> For example `35 % 10 = 5`. In our example `5` just so happens to be the last digit of the number.

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

Now that we have found a method to extract the last digit of a number accurately, we have a few more questions to address. How do we extract all its digits not just that the last digit? How do we add all the digits?

To address these questions, we need to simplify the approach to the problem. For now, we will only focus on 2-digit numbers (10 <= num <= 100). This makes it easier for us, as we only need to extract two digits then sum them. We already know how to extract the last digit, how then do we extract the first digit. One way to do this is to find a way to eliminate the digit we just extracted from `num` and then perform a modulo operation on the result to 