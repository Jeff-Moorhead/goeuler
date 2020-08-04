package main

import (
    "fmt"
    "os"
    "strconv"
)


func SieveOfEratosthenes(max int) []int {
    ints := make([]bool, max)  // create a slice of `max` elements

    for i := 2; i < max; i++ {
        ints[i] = true
    }

    for p := 2; p*p <= max; p++ {
        if ints[p] == true {
            for i := p * 2; i < max; i += p {
                ints[i] = false
            }
        }
    }

    var primes []int
    for i, v := range ints {
        if v == true {
            primes = append(primes, i)
        }
    }
    return primes
}


func PrimeFactors(max int) (factors []int) {
    primes := SieveOfEratosthenes(max)
    for _, v := range primes {
        if max % v == 0 {
            factors = append(factors, v)
        }
    }
    return
}


func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: largestprimef.go [NUM]")
        fmt.Println("NUM: The number to factorize.")
        return
    }

    max, err := strconv.Atoi(os.Args[1])

    if err != nil {
        fmt.Println("First argument to largestprime.go must be an integer")
        return
    }

    primes := PrimeFactors(max)
    fmt.Println(primes)
}
