package main

import "fmt"

func func1() {
	fmt.Println("Func 1")
}

func func2() {
	fmt.Println("Func 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada")
	fmt.Println("Está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// defer func1() // defer = adiar para antes do ultimo return  (se houver)
	// func2()

	fmt.Println(alunoAprovado(7, 4))
}
