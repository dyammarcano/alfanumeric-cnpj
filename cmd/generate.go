/*
Copyright © 2025 MadHouse madhouse@admin.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"github.com/dyammarcano/alfanumeric-cnpj/internal/cnpj"

	"github.com/spf13/cobra"
)

// generateCmd representa o comando generate
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gera um CNPJ alfanumérico válido",
	Long: `Gera um CNPJ alfanumérico válido, formata e valida o resultado.

Exemplo de uso:
  ./app generate

Este comando também mostra um exemplo de CNPJ inválido com DV alterado.`,
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generate() {
	// Gerando CNPJ válido
	valor := cnpj.GenerateCNPJ()
	fmt.Println("✅  CNPJ Gerado:", valor)

	// Formatando CNPJ
	fmt.Println("📎 CNPJ Formatado:", cnpj.FormatCNPJ(valor))

	// Validando CNPJ
	if cnpj.IsValid(valor) {
		fmt.Println("🔍 Validação: CNPJ gerado é válido ✅ ")
	} else {
		fmt.Println("🔍 Validação: CNPJ gerado é inválido ❌ ")
	}
}
