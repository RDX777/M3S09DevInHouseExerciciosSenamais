package main

import (
	"fmt"
	"math"
)

func encontraNumerosPrimos() {
	// Cria o array com 100 posições
	numeros := make([]int, 100)

	// Preenche o array com uma sequência numérica de 1 a 100
	for i := 0; i < 100; i++ {
		numeros[i] = i + 1
	}

	// Percorre o array e imprime os números primos
	for _, num := range numeros {
		if ehPrimo(num) {
			fmt.Print(num, " ")
		}
	}
}

func ehPrimo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	encontraNumerosPrimos()

}
