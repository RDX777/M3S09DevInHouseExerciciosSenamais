package main

import "fmt"

func main() {
	frutas := map[string]string{
		"maca":     "é uma fruta que cresce em arvore",
		"banana":   "é uma fruta amarela",
		"morango":  "O morango é uma fruta vermelha, cuja origem é a Europa",
		"melancia": "A melancia é uma fruta rasteira, originária da África",
	}

	fmt.Println("Descrição da maçã:", frutas["maca"])
	fmt.Println("Descrição da banana:", frutas["banana"])
	fmt.Println("Descrição do morango:", frutas["morango"])
	fmt.Println("Descrição da melancia:", frutas["melancia"])
}
