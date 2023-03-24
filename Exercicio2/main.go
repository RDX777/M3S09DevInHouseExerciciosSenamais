package main

import "fmt"

func main() {
    nums := []int{10, 5, 20, 15, 30}
    var max int

    for _, num := range nums {
        if num > max {
            max = num
        }
    }

    fmt.Printf("O maior valor existente nesse array é: %d\n", max)
}
