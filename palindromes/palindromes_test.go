package main

import "testing"


func TestIsPalindrome(t *testing.T) {
    palindrome := 252
    nonpalindrome := 253

    if ! IsPalindrome(palindrome) {
        t.Errorf("FAILURE: %v is a palindrome but returned false\n", palindrome)
    } else if IsPalindrome(nonpalindrome) {
        t.Errorf("FAILURE: %v is not a palindrome but returned true\n", nonpalindrome)
    }
}


func BenchmarkIsPalindrome(b *testing.B) {
    i := 1
    for i < 1000000 {
        IsPalindrome(i)
        i++
    }
}
