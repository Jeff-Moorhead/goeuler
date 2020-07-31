package main

import "fmt"


func Fibonacci(max int) []int {
    var fibs []int
    start, end := 1, 2
    for start <= max {
        fibs = append(fibs, start)
        start, end = end, start + end
    }
    return fibs
}


func Evenfibsum(max int) int {
    fibs := Fibonacci(max)
    sum := 0
    for _, f := range fibs {
        if f % 2 == 0 {
            sum += f
        }
    }
    return sum
}


func main() {
    sum := Evenfibsum(4000000)
    fmt.Println(sum)
}

