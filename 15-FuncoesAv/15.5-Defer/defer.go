package main

import "fmt"


func funcao1(){
	fmt.Println("Executando a Função 1")
}

func funcao2(){
	fmt.Println("Executando a Função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool{
	fmt.Println("Aluno")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}