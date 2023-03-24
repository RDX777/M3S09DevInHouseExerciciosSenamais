package main

import "fmt"

func main() {
	frutas := map[string]string{
		"maca":     "é uma fruta que cresce em arvore",
		"banana":   "é uma fruta amarela",
		"morango":  "O morango é uma fruta vermelha, cuja origem é a Europa",
		"melancia": "A melancia é uma fruta rasteira, originária da África",
	}

	if descricao, ok := frutas["morango"]; ok {
		fmt.Printf("morango - %s\n", descricao)
	} else {
		fmt.Println("Morango não encontrado no dicionário.")
	}
}
