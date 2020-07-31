package main

import "testing"


func TestFibonacci(t *testing.T) {
    expected := []int{1, 2, 3, 5, 8, 13, 21}
    got := Fibonacci(22)
    for i := range got {
        if got[i] != expected[i] {
            t.Errorf("FAILURE: got = %v, expected = %v", got, expected)
        }
    }
}


func TestEvenFibSum(t *testing.T) {
    expected := 10
    got := Evenfibsum(22)
    if got != expected {
        t.Errorf("FAILURE: got = %v, expected = %v", got, expected)
    }
}


func BenchmarkFibonaccie(t *testing.B) {
    for i := 5; i < 4000000; i++ {
        Fibonacci(i)
    }
}


func BenchmarkEvenFibSum(testing *testing.B) {
    for i := 5; i < 4000000; i++ {
        Evenfibsum(i)
    }
}

