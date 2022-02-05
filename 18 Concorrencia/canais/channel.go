package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo.", canal)

	fmt.Println("Depois da funcao escrever")

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim")
}

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}
	close(canal)
}
