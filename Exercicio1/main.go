package main

import "fmt"

func sumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	arr := []int{3, 5, 2}
	sum := sumArray(arr)
	fmt.Println(sum)
}
