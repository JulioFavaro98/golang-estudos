package enderecos

import "testing"

func TesteTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua ABC"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo recebido é diferente do esperado")
	}

}
