package main

import "fmt"

func concatenaStrings(str1 string, str2 string) string {
	return str1 + str2
}

func main() {
	var texto1 string
	var texto2 string
	var retorno string

	texto1 = "Texto 1 "
	texto2 = "Texto 2"

	retorno = concatenaStrings(texto1, texto2)

	fmt.Println(retorno)

}
