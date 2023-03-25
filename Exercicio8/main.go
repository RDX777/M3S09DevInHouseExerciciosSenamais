// go env -w GO111MODULE=off Usar para importar sem git

package main

import (
	"fmt"

	"./labmath"
)

func main() {
	var numero int
	numero = 7
	if labmath.Primo(numero) {
		fmt.Printf("%d é primo\n", numero)
	} else {
		fmt.Printf("%d não é primo\n", numero)
	}

}
