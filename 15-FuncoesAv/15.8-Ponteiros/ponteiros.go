package main

import "fmt"

func inverterSinal(numero *int){
	 *numero = *numero * -1

}

func main() {
	// numero := 20
	// numeroInvertido := inverterSinal(numero)
	// fmt.Println(numeroInvertido)


	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinal(&novoNumero)
	fmt.Println(novoNumero)

}
