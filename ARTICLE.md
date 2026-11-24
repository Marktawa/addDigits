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

The final solution should use a function called `addDigits`. 

Create a file named `main.go` and add the following code:
```go
package main

import "fmt"

func addDigits(num int) int {
    return num
}

func main() {
    fmt.Println(addDigits(9))
    fmt.Println(addDigits(3))
    fmt.Println(addDigits(0))
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
```

The program returns the value of the integer, each time `addDigits` is called. No computation is done.

We can set conditions that if a number is between 0 and 9 inclusive then the function should just return the same number as it is already a single digit. For any other number we will set it to return `-1` for now.

Update `main.go` with the following code:
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
    fmt.Println(addDigits(38))
}
```

Save the file and run it in your terminal:
```shell
go run main.go
```

You should see the following output in your terminal:
```
9
3
0
-1
```

## Two Digit Numbers

Now that we have handled single digit numbers, we can move on.

If a number has more than one digit, we need to first find a way to extract the digits of an integer. One way to do this is by making use of the `%` (pronounced modulo) operator. This gives the remainder when you divide a number by another number. For example, `5 / 2 = 2 remainder 1` therefore `5 % 2 = 1`.

If you find the modulo when a number is divided by 10, the answer you get is always the last digit of that number. <!-- Look for supporting theorems --> For example `38 % 10 = 8`. In our example `8` is the last digit of the number.

Before we proceed, let's test this out in Go. 

Create a file named `mod10.go` and add the following code:
```go
package main

import "fmt"

func main() {
    fmt.Println(38 % 10)
}
```

Save the file and run the code in your shell:
```shell
go run mod10.go
```

You should see `8` printed in your shell. 

<!-- The final solution should use a function called `addDigits`. So here's how it would look  like using that: -->

We can test out using a function too. As seen in the code below. Update `mod10.go`:

```go
package main

import "fmt"

func mod10(num int) int {
    return num % 10
}

func main() {
    fmt.Println(mod10(38))
}
```

You should see `8` printed in your shell like in the previous code snippet.

Now that we have found a method to extract the last digit of a number accurately, we have a few more questions to address. How do we extract all its digits not just that the last digit? How do we add all the digits?

To address these questions, we need to simplify the approach to the problem. For now, we will only focus on 2-digit numbers (10 <= num <= 100). This makes it easier for us, as we only need to extract two digits then sum them. We already know how to extract the last digit, how then do we extract the first digit? 

To extract the next digit, just divide the number by 10.

For our example number 38, `38 / 10 = 3`

In integer division there are no decimal points and Go programming language floors the decimal value, hence we can confidently get the first digit simply by dividing the number by 10.

<!-- First, we need to remove the digit we just extracted.

For example for the number 35, if we have identified 5 as the last digit, we need to remove from 35 so that it becomes 3. This will help us identify 3 as the next digit -->

Create a new file named `div10.go`:
```shell
touch div10.go
```

This will contain a simple program to show the results of dividing by 10.

Add the following code to `div10.go`:
```go
package main

import "fmt"

func div10(num int) int {
    return num / 10
}

func main() {
    fmt.Println(div10(38))
}
```

Save `div10.go` and run the file to test the code:
```shell
go run div10.go
```

You should see the following output:
```
3
```

Integer division as witnessed. Proof that Go floors(rounds down) the decimal point.

<!-- 
Add sources that people can use to get up to speed with Go and follow along with the tutorial. Prerequisites like basic programming knowledge, math and logic knowledge are required to follow along. Go topics for people to catch should be picked from the Go tour website. It's hands on and concise. Good enough to prime readers 
-->

Now that we have found a way to extract the first digit and the last digit of a two-digit number, we then have to sum the two. Check if the sum is a two-digit number. If the sum is a single digit number, it can be returned and if it is a two-digit number, we have to extract the number again.

You maybe wondering if this process of adding the digits to create the sum then checking if the sum is more than 9 is infinite, as if there's a set of 2-digit numbers that when added will continue to produce 2-digit numbers. Well the answer is that no, such a set of numbers doesn't exist.

<!-- Prompts should comment out or strike through edited statements and return the markdown copy -->

Here's the result of adding 2-digit numbers:

```
10 --> 1 + 0 --> 1
11 --> 1 + 1 --> 2
12 --> 1 + 2 --> 3
13 --> 1 + 3 --> 4
14 --> 1 + 4 --> 5
15 --> 1 + 5 --> 6
16 --> 1 + 6 --> 7
17 --> 1 + 7 --> 8
18 --> 1 + 8 --> 9
19 --> 1 + 9 --> 10, 10 --> 1 + 0 --> 1
20 --> 2 + 0 --> 2
21 --> 2 + 1 --> 3
22 --> 2 + 2 --> 4
23 --> 2 + 3 --> 5
24 --> 2 + 4 --> 6
25 --> 2 + 5 --> 7
26 --> 2 + 6 --> 8
27 --> 2 + 7 --> 9
28 --> 2 + 8 --> 10, 10 --> 1 + 0 --> 1
29 --> 2 + 9 --> 11, 11 --> 1 + 1 --> 2
30 --> 3 + 0 --> 3
31 --> 3 + 1 --> 4
32 --> 3 + 2 --> 5
33 --> 3 + 3 --> 6
34 --> 3 + 4 --> 7
35 --> 3 + 5 --> 8
36 --> 3 + 6 --> 9
37 --> 3 + 7 --> 10, 10 --> 1 + 0 --> 1
38 --> 3 + 8 --> 11, 11 --> 1 + 1 --> 2
39 --> 3 + 9 --> 12, 12 --> 1 + 2 --> 3
40 --> 4 + 0 --> 4
41 --> 4 + 1 --> 5
42 --> 4 + 2 --> 6
43 --> 4 + 3 --> 7
44 --> 4 + 4 --> 8
45 --> 4 + 5 --> 9
46 --> 4 + 6 --> 10, 10 --> 1 + 0 --> 1
47 --> 4 + 7 --> 11, 11 --> 1 + 1 --> 2
48 --> 4 + 8 --> 12, 12 --> 1 + 2 --> 3
49 --> 4 + 9 --> 13, 13 --> 1 + 3 --> 4
50 --> 5 + 0 --> 5
51 --> 5 + 1 --> 6
52 --> 5 + 2 --> 7
53 --> 5 + 3 --> 8
54 --> 5 + 4 --> 9
55 --> 5 + 5 --> 10, 10 --> 1 + 0 --> 1
56 --> 5 + 6 --> 11, 11 --> 1 + 1 --> 2
57 --> 5 + 7 --> 12, 12 --> 1 + 2 --> 3
58 --> 5 + 8 --> 13, 13 --> 1 + 3 --> 4
59 --> 5 + 9 --> 14, 14 --> 1 + 4 --> 5
60 --> 6 + 0 --> 6
61 --> 6 + 1 --> 7
62 --> 6 + 2 --> 8
63 --> 6 + 3 --> 9
64 --> 6 + 4 --> 10, 10 --> 1 + 0 --> 1
65 --> 6 + 5 --> 11, 11 --> 1 + 1 --> 2
66 --> 6 + 6 --> 12, 12 --> 1 + 2 --> 3
67 --> 6 + 7 --> 13, 13 --> 1 + 3 --> 4
68 --> 6 + 8 --> 14, 14 --> 1 + 4 --> 5
69 --> 6 + 9 --> 15, 15 --> 1 + 5 --> 6
70 --> 7 + 0 --> 7
71 --> 7 + 1 --> 8
72 --> 7 + 2 --> 9
73 --> 7 + 3 --> 10, 10 --> 1 + 0 --> 1
74 --> 7 + 4 --> 11, 11 --> 1 + 1 --> 2
75 --> 7 + 5 --> 12, 12 --> 1 + 2 --> 3
76 --> 7 + 6 --> 13, 13 --> 1 + 3 --> 4
77 --> 7 + 7 --> 14, 14 --> 1 + 4 --> 5
78 --> 7 + 8 --> 15, 15 --> 1 + 5 --> 6
79 --> 7 + 9 --> 16, 16 --> 1 + 6 --> 7
80 --> 8 + 0 --> 8
81 --> 8 + 1 --> 9
82 --> 8 + 2 --> 10, 10 --> 1 + 0 --> 1
83 --> 8 + 3 --> 11, 11 --> 1 + 1 --> 2
84 --> 8 + 4 --> 12, 12 --> 1 + 2 --> 3
85 --> 8 + 5 --> 13, 13 --> 1 + 3 --> 4
86 --> 8 + 6 --> 14, 14 --> 1 + 4 --> 5
87 --> 8 + 7 --> 15, 15 --> 1 + 5 --> 6
88 --> 8 + 8 --> 16, 16 --> 1 + 6 --> 7
89 --> 8 + 9 --> 17, 17 --> 1 + 7 --> 8
90 --> 9 + 0 --> 9
91 --> 9 + 1 --> 10, 10 --> 1 + 0 --> 1
92 --> 9 + 2 --> 11, 11 --> 1 + 1 --> 2
93 --> 9 + 3 --> 12, 12 --> 1 + 2 --> 3
94 --> 9 + 4 --> 13, 13 --> 1 + 3 --> 4
95 --> 9 + 5 --> 14, 14 --> 1 + 4 --> 5
96 --> 9 + 6 --> 15, 15 --> 1 + 5 --> 6
97 --> 9 + 7 --> 16, 16 --> 1 + 6 --> 7
98 --> 9 + 8 --> 17, 17 --> 1 + 7 --> 8
99 --> 9 + 9 --> 18, 18 --> 1 + 8 --> 9
```

From this we can deduce that all 2-digit numbers only have to be added at most twice to produce a single digit number. Therefore checking whether the result of adding digits of a two digit number produces a single digit only needs to be done once.

Given this discovery, the process becomes:
```
- extract digits
- sum the digits
- check if sum is a double digit
    - if it is a double digit number, extract the digits
    - sum the digits
- return sum
```

We can then update `main.go` with the following code:
```go
package main

import "fmt"

func addDigits(num int) int {
    var digit1 int = num % 10
    var digit2 int = num / 10

    var sum int = digit1 + digit2
    if sum > 9 {
        digit1 = sum % 10
        digit2 = sum / 10

        sum = digit1 + digit2
    }

    return sum
}

func main() {
    fmt.Println(addDigits(0))
    fmt.Println(addDigits(38))
}
```

Run the code:
```shell
go run main.go
```

You should see the following output:
```
0
2
```

**Explanation**: The process is
38 --> 3 + 8 --> 11, 11 --> 1 + 1 --> 2

- `digit1` represents the last digit of a 2-digit number
- `digit2` is the first digit
- `sum` is the result of adding `digit1` and `digit2`
- If `sum` is greater than 9(highest single digit number) then we need to restart the process
