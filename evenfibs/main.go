package main

import "fmt"


func fibonacci(max int) []int {
    var fibs []int
    start, end := 1, 2
    for start <= max {
        fibs = append(fibs, start)
        start, end = end, start + end
    }
    return fibs
}


func evenfibsum(max int) int {
    fibs := fibonacci(max)
    sum := 0
    for _, f := range fibs {
        if f % 2 == 0 {
            sum += f
        }
    }
    return sum
}


func main() {
    sum := evenfibsum(4000000)
    fmt.Println(sum)
}

