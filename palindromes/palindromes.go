package main

import (
    "fmt"
)


func IsPalindrome(n int) bool {
    original := n
    reverse := 0
    for n > 0 {
        remainder := n % 10
        reverse = (reverse * 10) + remainder
        n /= 10
    }

    return original == reverse
}


func main() {
    palindrome := 0
    for i := 999; i > 0; i-- {
        for j := 999; j > 0; j-- {
            product := i * j

            if product < palindrome {
                // No need to continue the inner loop,
                // Every other product will be smaller
                break
            }

            if IsPalindrome(product) {
                palindrome = product
            }
        }
    }
    fmt.Printf("The greatest palindromic product of two three-digit numbers is %v.\n", palindrome)
}
