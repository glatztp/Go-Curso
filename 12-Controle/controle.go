package main

import "fmt"

func main() {
	fmt.Println("Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if numero2 := numero; numero2 > 0 {
		fmt.Println("Numero 2 é maior que 0")
	} else {
		fmt.Println("Numero 2 é menor que 0")
	}
}
