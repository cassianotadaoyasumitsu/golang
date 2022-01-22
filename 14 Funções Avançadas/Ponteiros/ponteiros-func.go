package main

import "fmt"

func inverterSinal(n int) int {
	return n * -1
}

func inverterSinalComPonteiro(n *int) {
	*n = *n * -1
}

func main() {
	n := 20
	numeroInvertido := inverterSinal(n)
	fmt.Println(numeroInvertido)
	fmt.Println(n)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
