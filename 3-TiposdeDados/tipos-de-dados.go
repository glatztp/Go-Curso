package main

import (
	"errors"
	"fmt"
)

func main(){
	var numero int = 100
	fmt.Println(numero)

	var numeroReal float32 =1223.23
	fmt.Println(numeroReal)

	var str string = "texto"
	fmt.Println(str)

	var texto string
	fmt.Println(texto)

	var boolean bool = true
	fmt.Println(boolean)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}