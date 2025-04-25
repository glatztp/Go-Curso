package main

import "fmt"

func recuperarExe(){
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada co sucesso")
	}

}
func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExe()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media e exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")

}