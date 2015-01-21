package main

import "fmt"

func main() {
    largest_palindrome := 0
    i := 999
    for i >= 100 {
        j := 999
        for j >= i {
            test := i * j
            if is_palindrome(test) && test > largest_palindrome {
                largest_palindrome = test
            }
            j--
        }
        i--
    }
    fmt.Println(largest_palindrome)
}

func reverse(n int) int{
    reversed := 0
    for n != 0 {
        reversed = reversed * 10 + n % 10
        n = n / 10
    }
    return reversed
}


func is_palindrome(n int) bool {
    return n == reverse(n)
}
