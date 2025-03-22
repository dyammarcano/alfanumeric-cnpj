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

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate [CNPJ...]",
	Short: "Valida um ou mais CNPJs alfanuméricos",
	Long: `Valida um ou mais CNPJs alfanuméricos, com ou sem máscara.

Exemplos de uso:
  ./app validate OT.WXQ.ENJ/DKC6-20
  ./app validate RZ.YYO.MTN/OLSV-26 VX7VLX1I5M4X05 RZYYOMTNOLSV26 JJQFNXSNR8FD58 VX.7VL.X1I/5M4X-05`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("⚠️  Nenhum CNPJ foi informado. Por favor, passe pelo menos um argumento para validação.")
			return
		}
		for i, valor := range args {
			if cnpj.IsValid(valor) {
				cmd.Printf("[%d] ✅  CNPJ válido:   %s\n", i+1, cnpj.FormatCNPJ(valor))
			} else {
				cmd.Printf("[%d] ❌  CNPJ inválido: %s\n", i+1, cnpj.FormatCNPJ(valor))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
