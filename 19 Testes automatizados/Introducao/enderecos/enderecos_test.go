package enderecos // enderecos_test (importar package e chamar a funcao endereco.TipoDeEndereco)

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua abc", "Rua"},
		{"Avenida Pualista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Invalido"},
		{"Estrada qualquer", "Estrada"},
		{"RUA dos bobos", "Rua"},
		{"AVENIDA das ruas", "Avenida"},
		{"", "Tipo Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel() // Roda testes em paralelo
	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}

// go test --coverprofile {nomedoarquivo.txt} "cria arquivo"
// go tool cover --html={nomedoarquivo.txt} "le arquivo"
