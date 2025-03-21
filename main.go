package main

import (
	"AlfanumericCNPJ/internal/cnpjalfanumerico"
	"fmt"
)

func main() {
	// Gerando CNPJ válido
	cnpj := cnpjalfanumerico.GenerateCNPJ()
	fmt.Println("CNPJ Gerado:", cnpj)

	// Formatando CNPJ
	fmt.Println("CNPJ Formatado:", cnpjalfanumerico.FormatCNPJ(cnpj))

	// Validando CNPJ
	fmt.Println("CNPJ é válido?", cnpjalfanumerico.ValidateCNPJ(cnpj))

	// Teste com um CNPJ inválido
	cnpjInvalido := cnpj[:len(cnpj)-2] + "99"
	fmt.Println("CNPJ Inválido:", cnpjInvalido, "->", cnpjalfanumerico.ValidateCNPJ(cnpjInvalido))
}
