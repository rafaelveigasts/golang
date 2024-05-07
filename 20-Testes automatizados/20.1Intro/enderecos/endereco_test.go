// teste unitário

package enderecos_test

import (
	. "testes/enderecos"
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
    testCases := []struct {
        name     string
        endereco string
        want     string
    }{
        {"Rua case", "Rua das Flores", "rua"},
        {"Avenida case", "Avenida Paulista", "avenida"},
        {"Estrada case", "Estrada de Santos", "estrada"},
        {"Rodovia case", "Rodovia dos Imigrantes", "rodovia"},
        // {"Praça case", "Praça 14 bis", "praça"},
        {"Invalid type case", "Alameda das Rosas", "Tipo Inválido"},
        {"Empty string case", "", "Tipo Inválido"},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            got := TipoDeEndereco(tc.endereco)
            if got != tc.want {
                t.Errorf("TipoDeEndereco(%q) = %q; want %q", tc.endereco, got, tc.want)
            }
        })
    }
}

// para rodar o teste independente do diretório que está, basta rodar o comando: go test ./...
// para rodar o teste de um pacote específico, basta rodar o comando: go test ./enderecos
// para rodar o teste de um arquivo específico, basta rodar o comando: go test endereco_test.go
// para rodar o teste de um arquivo específico e ver o log de execução, basta rodar o comando: go test -v endereco_test.go
// para rodar o teste de um arquivo específico e ver o log de execução e o tempo de execução, basta rodar o comando: go test -v -bench . endereco_test.go
// para rodar o teste de um arquivo específico e ver o log de execução e o tempo de execução, basta rodar o comando: go test -v -bench . -benchtime 10s endereco_test.go
// para ver o coverage do teste, basta rodar o comando: go test -cover ./enderecos
// para ver o coverage do teste e gerar um arquivo com o resultado, basta rodar o comando: go test -coverprofile cobertura.txt ./enderecos
// para ver o coverage do teste e gerar um arquivo com o resultado e abrir o arquivo no navegador, basta rodar o comando: go test -coverprofile cobertura.txt ./enderecos && go tool cover --html=cobertura.txt