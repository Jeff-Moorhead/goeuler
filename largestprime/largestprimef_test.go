package main

import "testing"


func TestSieveOfEratosthenes(t *testing.T) {
    lt10 := []int{2, 3, 5, 7}
    got := SieveOfEratosthenes(10)
    if len(got) != len(lt10) {
        t.Errorf("FAILURE: Slices are different lengths: expected %v, got %v\n", len(lt10), len(got))
        return
    }

    for i, v := range got {
        if v != lt10[i] {
            t.Errorf("FAILURE: Incorrect value at i = %v, v = %v\n", i, v)
        }
    }
}


func TestPrimeFactors(t *testing.T) {
    factors30 := []int{2, 3, 5}
    got := PrimeFactors(30)

    if len(got) != len(factors30) {
        t.Errorf("FAILURE: different lengths: expected %v, got %v\n", len(factors30), len(got))
        return
    }

    for i, v := range got {
        if v != factors30[i] {
            t.Errorf("FAILURE: Incorrect value at i = %v, v = %v. Should be %v\n", i, v, factors30[i])
        }
    }
}
