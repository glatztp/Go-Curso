package main

import "fmt"

var n int 

func init(){
	fmt.Println("Executando a funcao main")
	n = 10

}

func main(){
	fmt.Println("Funcao main sendo executada")
	fmt.Println(n)
}