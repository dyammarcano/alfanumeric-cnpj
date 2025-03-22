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
