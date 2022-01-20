package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá Mundo")
	}()

	func(text string) {
		fmt.Println(text)
	}("Cassiano")

	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Passando parametro")

	fmt.Println(retorno)
}
