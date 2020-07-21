package main

import (
    "testing"
    "math"
)


func TestTriangularSum(t *testing.T) {
    res := map[int]int{
        1: 1,
        2: 3,
        3: 6,
        4: 10,
    }
    for i := 1; i < 5; i++ {
        got := TriangularSum(1, i)

        if got != res[i] {
            t.Errorf("FAILURE: i = %d, got = %d, expected %d", i, got, res[i])
        }
    }
}


func TestGetMultiples(t *testing.T) {
    res := map[int]int {
        1: 0,
        2: 0,
        3: 0,
        4: 3,
        5: 3,
        6: 8,
        7: 14,
        8: 14,
        9: 14,
        10: 23,
    }
    for i := 1; i <= 10; i++ {
        got := getMultiples(i)

        if got != res[i] {
            t.Errorf("FAILURE: i = %d, got = %d, expected %d", i, got, res[i])
        }
    }
}


func BenchmarkTriangularSum(b *testing.B) {
    for i := 0; i < 10; i++ {
        num := int(math.Pow10(i))
        TriangularSum(1, num)
    }
}


func BenchmarkGetMultiples(b *testing.B) {
    for i := 0; i < 10; i++ {
        num := int(math.Pow10(i))
        getMultiples(num)
    }
}

