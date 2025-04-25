package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)

	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Davi", "Lucas"}
	
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	

	usuario := map[string]string {
		"nome":					"Leo",
		"sobrenome": 		"Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
