package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoinserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		// {"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoinserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O Tipo recebido %s e diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido é diferente do esperado! \n Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	// }

}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste Quebrou!")
	}
}

// func TestTipoDeEndereco(t *testing.T) {
// 	enderecoParaTeste := "Rua sao jose"

// 	tipoDeEnderecoEsperado := "Avenida"

// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 		t.Errorf("O tipo recebido é diferente do esperado! \n Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
// 	}

// }

// go test --cover
// go test -ver
