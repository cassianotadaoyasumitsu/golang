package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	time.Sleep(time.Second)
	// 	fmt.Println("i++")
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("j++", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := []string{"John", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Cassiano",
		"sobrenome": "Yasumitsu",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Loop Infinito")
	}
}
