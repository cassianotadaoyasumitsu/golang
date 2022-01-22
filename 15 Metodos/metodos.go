package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando usuario %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Cassiano", 42}
	usuario2.salvar()

	maior := usuario2.maiorIdade()
	fmt.Println(maior)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
