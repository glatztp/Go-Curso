package main

import "fmt"


type pessoa struct {
	nome 			string
	sobrenome string
	idade 		uint8
	altura 		uint8

}

type estudante struct {
	pessoa
	curso 		string
	faculdade string

}


func main() {
	fmt.Printf("Heranca")

	p1 := pessoa{"Joao", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Go", "Uniaselvi"}
	fmt.Println(e1)

}