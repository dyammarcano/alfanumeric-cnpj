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
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dyammarcano/alfanumeric-cnpj/pkg/cnpj"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
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

var (
	pgHost     string
	pgPort     int
	pgUser     string
	pgPassword string
	pgDatabase string
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Inicia um servidor HTTP para validar e formatar CNPJs via JSON",
	Long: `Este comando inicia um servidor HTTP usando Echo que escuta requisições POST com JSON
para validar, formatar e calcular DV de CNPJs alfanuméricos.

Exemplo de chamada com curl:
curl -X POST http://localhost:4400/api/cnpj/validate -H "Content-Type: application/json" -d '{"cnpj":"GIFZXOWDNZYM58"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		if pgHost == "" || pgUser == "" || pgPassword == "" || pgDatabase == "" {
			cmd.Println("Erro: é necessário informar --pg-host, --pg-user, --pg-password e --pg-database")
			cmd.Usage()
			os.Exit(1)
		}

		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			pgHost, pgPort, pgUser, pgPassword, pgDatabase)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("Erro ao conectar no banco:", err)
		}
		defer db.Close()

		if err := db.Ping(); err != nil {
			log.Fatal("Não foi possível pingar o banco:", err)
		}

		// Crear tabla si no existe
		if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS cnpjs (id SERIAL PRIMARY KEY,cnpj TEXT NOT NULL UNIQUE);`); err != nil {
			log.Fatal(err)
		}

		http.HandleFunc("GET /api/cnpj/generate", generateHandler(db))
		http.HandleFunc("POST /api/cnpj/validate", validateHandler)

		log.Println("🚀 Servidor iniciado em http://localhost:4400")
		log.Fatal(http.ListenAndServe(":4400", nil))
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVar(&pgHost, "pg-host", "", "PostgreSQL host")
	apiCmd.Flags().IntVar(&pgPort, "pg-port", 5432, "PostgreSQL port")
	apiCmd.Flags().StringVar(&pgUser, "pg-user", "", "PostgreSQL user")
	apiCmd.Flags().StringVar(&pgPassword, "pg-password", "", "PostgreSQL password")
	apiCmd.Flags().StringVar(&pgDatabase, "pg-database", "", "PostgreSQL database")
}

func generateHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var (
			cnpjValue string
			formatado string
			valido    bool
			dv        string
			exists    bool
		)

		maxAttempts := 100
		for attempts := 0; attempts < maxAttempts; attempts++ {
			cnpjValue = cnpj.GenerateCNPJ()

			if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM cnpjs WHERE cnpj = $1)", cnpjValue).Scan(&exists); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(CNPJResponse{Erro: "erro ao consultar o banco"})
				return
			}

			if !exists {
				if _, err := db.Exec("INSERT INTO cnpjs (cnpj) VALUES ($1)", cnpjValue); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(w).Encode(CNPJResponse{Erro: "erro ao salvar no banco"})
					return
				}
				break
			}
		}

		if exists {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(CNPJResponse{Erro: "não foi possível gerar um CNPJ único após várias tentativas"})
			return
		}

		formatado = cnpj.FormatCNPJ(cnpjValue)
		valido = cnpj.IsValid(cnpjValue)
		dv, _ = cnpj.CalculateDV(cnpjValue)

		_ = json.NewEncoder(w).Encode(CNPJResponse{
			CNPJOriginal: cnpjValue,
			Formatado:    formatado,
			Valido:       valido,
			DV:           dv,
		})
	}
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req CNPJRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CNPJResponse{Erro: "request inválido"})
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CNPJResponse{Erro: "JSON inválido"})
		return
	}

	newCNPJResponse := NewCNPJResponse(req.CNPJ)
	if newCNPJResponse.Valido {
		newCNPJResponse.DV, err = cnpj.CalculateDV(newCNPJResponse.CNPJOriginal)
		if err != nil {
			newCNPJResponse.Erro = err.Error()
		}
	}

	json.NewEncoder(w).Encode(newCNPJResponse)
}
