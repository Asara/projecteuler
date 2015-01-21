package main

import "fmt"

func main() {
    if (is_palindrome(9009)) {
        fmt.Println("Yes")
    }

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
