package main

import (
    "fmt"
    "os"
    "strconv"
)


const usage string = "Usage: mults3and5 [MAX] (int)\n" +
"Positional arguments: MAX - The upper bound for summing (e.g. 1000 for Project Euler #1).\n" +
"Multiples of 3 and 5 strictly less than MAX will be included in the end result."


/* Brute force folution: O(n) */

func getMultiples(max int) int {
    sum := 0
    for i := 0; i < max; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }
    return sum
}


func TriangularSum(base int, max int) int {
    /************************************************************** 

    base: The multiplier who's multiples we want to sum (3 or 5 for
    Project Euler #1
    
    max: The upper limit of multiples (1000 for PE #1)

    Uses the triangular numbers formula to reduce runtime to O(1) 
    Triangular number formula: 

                              n(n+1)
    Sum from k=1 to n of k =  ------
                                2
    ***************************************************************/
    numerator := (max/base)*((max/base)+1)
    return base * (numerator / 2)
}


func main() {
    if len(os.Args) < 2 {
        fmt.Println(usage)
        return
    }

    max, err := strconv.Atoi(os.Args[1])
    max -= 1 // Subtract 1 to add multiples strictly less than max

    if err != nil {
        fmt.Println(usage)
        return
    }
    sum := TriangularSum(3, max) + TriangularSum(5, max) - TriangularSum(15, max)
    fmt.Println(sum)
}

