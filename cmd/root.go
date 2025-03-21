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
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa o comando base quando nenhum subcomando é fornecido
var rootCmd = &cobra.Command{
	Use:   "AlfanumericCNPJ",
	Short: "Ferramenta CLI para gerar, validar e formatar CNPJs alfanuméricos",
	Long: `📦 AlfanumericCNPJ é uma ferramenta de linha de comando para gerar, validar e formatar CNPJs compostos por letras e números.

Comandos disponíveis:
  • generate  → Gera um novo CNPJ válido
  • validate  → Valida um ou mais CNPJs fornecidos
  • format    → Aplica a máscara padrão em CNPJs alfanuméricos

Exemplo de uso:
  ./AlfanumericCNPJ generate
  ./AlfanumericCNPJ validate GI.FZX.OWD/NZYM-40
  ./AlfanumericCNPJ format ABCDEFGHIJKL80`,
}

// Execute executa o comando root e todos os subcomandos registrados
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
