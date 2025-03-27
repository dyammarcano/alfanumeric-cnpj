# AlfanumericCNPJ

📦 Uma ferramenta de linha de comando em Go para **gerar**, **validar** e **formatar** CNPJs alfanuméricos
personalizados.

---

## ✨ Funcionalidades

- ✅ Geração de CNPJs alfanuméricos válidos (com letras e números)
- 🔍 Validação de CNPJs com ou sem máscara
- 🔠 Formatação no padrão `##.###.###/####-##`
- ⚠️ Detecção de CNPJs inválidos e DV incorretos
- 📦 Estruturado com [Cobra CLI](https://github.com/spf13/cobra)

---

## 📦 Instalação

```bash
git clone https://github.com/dyammarcano/alfanumeric-cnpj.git
cd alfanumeric-cnpj
go build -o app
```

## Run server
```bash
podman run -d --name cnpj-postgres -e POSTGRES_USER=cnpjuser -e POSTGRES_PASSWORD=cnpjpass -e POSTGRES_DB=cnpjdb -p 5432:5432 -v pgdata:/var/lib/postgresql/data docker.io/postgres:15
app api --pg-host=localhost --pg-port=5432 --pg-user=cnpjuser --pg-password=cnpjpass --pg-database=cnpjdb
```

## Use in your code
```go
package main

import (
    "fmt"
    "github.com/dyammarcano/alfanumeric-cnpj/pkg/cnpj"
)

func main() {
    value := cnpj.FormatCNPJ("OTWXQENJDKC620")
    fmt.Printf("CNPJ: %s, valid: %v",value,cnpj.IsValid(value))
}
```
