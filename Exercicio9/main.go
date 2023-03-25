package main

import (
	"fmt"
)

func parOuImpar(num int) string {
	if num%2 == 0 {
		return "par"
	} else {
		return "ímpar"
	}
}

func main() {
	fmt.Println(parOuImpar(8))

}
