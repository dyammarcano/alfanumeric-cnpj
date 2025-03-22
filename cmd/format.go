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
	"github.com/dyammarcano/alfanumeric-cnpj/internal/cnpj"

	"github.com/spf13/cobra"
)

// formatCmd representa o comando 'format'
var formatCmd = &cobra.Command{
	Use:   "format [CNPJ...]",
	Short: "Formata um ou mais CNPJs alfanuméricos",
	Long: `Formata CNPJs no padrão ##.###.###/####-##, mesmo que estejam sem máscara.

Exemplos de uso:
  ./app format OTWXQENJDKC620
  ./app format RZYYOMTNOLSV26 D6RJ1CUTQQAA22`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("⚠️  Nenhum CNPJ foi informado. Informe pelo menos um valor para formatar.")
			return
		}

		for i, valor := range args {
			formatado := cnpj.FormatCNPJ(valor)
			cmd.Printf("[%d] 🧾 Original:  %s\n    📎 Formatado: %s\n", i+1, valor, formatado)
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
