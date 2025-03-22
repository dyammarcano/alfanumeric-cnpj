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
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"net/http"
)

// CNPJRequest estrutura de requisição esperada
type CNPJRequest struct {
	CNPJ string `json:"cnpj"`
}

// CNPJResponse estrutura de resposta
type CNPJResponse struct {
	CNPJOriginal string `json:"cnpj_original"`
	Formatado    string `json:"formatado"`
	DV           string `json:"dv,omitempty"`
	Valido       bool   `json:"valido"`
	Erro         string `json:"erro,omitempty"`
}

func NewCNPJResponse(value string) *CNPJResponse {
	return &CNPJResponse{
		CNPJOriginal: value,
		Formatado:    cnpj.FormatCNPJ(value),
		Valido:       cnpj.IsValid(value),
	}
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Inicia um servidor HTTP para validar e formatar CNPJs via JSON",
	Long: `Este comando inicia um servidor HTTP usando Echo que escuta requisições POST com JSON
para validar, formatar e calcular DV de CNPJs alfanuméricos.

Exemplo de chamada com curl:
curl -X POST http://localhost:4400/api/cnpj -H "Content-Type: application/json" -d '{"cnpj":"GIFZXOWDNZYM58"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		e.POST("/api/cnpj", func(c echo.Context) error {
			var req CNPJRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, CNPJResponse{Erro: "JSON inválido"})
			}

			newCNPJResponse := NewCNPJResponse(req.CNPJ)
			var err error
			newCNPJResponse.DV, err = cnpj.CalculateDV(newCNPJResponse.CNPJOriginal)
			if err != nil {
				newCNPJResponse.Erro = err.Error()
			}

			return c.JSON(http.StatusOK, newCNPJResponse)
		})

		e.Logger.Fatal(e.Start(":4400"))
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
