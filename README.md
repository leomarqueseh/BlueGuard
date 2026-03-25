# 🔐 BlueGuard

BlueGuard é uma ferramenta de segurança ofensiva (Web Security Scanner) desenvolvida em Go, com foco em detecção e validação real de vulnerabilidades (PoC), inspirada em ferramentas como Nessus.

---

## 🚀 Features atuais

- 🔍 Scan de aplicações web
- ⚙️ Engine modular baseada em plugins
- 🧠 Risk Engine (pontuação de vulnerabilidades)
- 🌐 Dashboard Web (estilo SaaS)
- 📄 Output em JSON
- 🔎 Fingerprinting de tecnologias
- 🛡️ Header Analysis

---

## 🔥 Vulnerabilidades suportadas

### ✅ Open Redirect (PoC real)
- Teste ativo com payload
- Validação via header `Location`
- Sem falso positivo

### ✅ Git Exposed (CONFIRMED)
- Detecta:
  - `.git/HEAD`
  - `.git/config`
- Validação por conteúdo real

### 🔥 Git Dump (CRITICAL)
- Baixa arquivos sensíveis do repositório
- Possível reconstrução do código fonte

---

## ⚠️ Aviso Legal

Esta ferramenta deve ser utilizada apenas em:

- Ambientes próprios
- Laboratórios de teste
- Programas de Bug Bounty autorizados

Uso indevido pode violar leis como a LGPD e legislações de crimes digitais.

---

## 🖥️ Instalação

```bash
git clone https://github.com/leomarqueseh/BlueGuard.git
cd BlueGuard
go mod tidy
go build -o blueguard ./cmd/blueguard
````

---

## ▶️ Uso

### Scan simples

```bash
./blueguard -u https://example.com
```

### Lista de alvos

```bash
./blueguard -l targets.txt
```

### Dashboard Web

```bash
./blueguard -web
```

Acesse:

```
http://localhost:8080
```

---

## 📊 Exemplo de saída

```json
[
  {
    "title": "Git Repository Exposed (CONFIRMED)",
    "severity": "HIGH",
    "target": "https://target/.git/config",
    "score": 9.5
  }
]
```

---

## 🧠 Roadmap

* [ ] Git full dump (reconstrução completa)
* [ ] XSS detection (PoC)
* [ ] SSRF detection
* [ ] Directory Listing
* [ ] Login system (SaaS)
* [ ] API REST
* [ ] Histórico de scans (SQLite)
* [ ] Export PDF

---

## 👨‍💻 Autor

Leonardo Matheus Marques Ferreira

---

## ⭐ Contribuição

Pull requests são bem-vindos!

````

---

# 🚀 AGORA: GIT DUMP COMPLETO (NÍVEL MONSTRO)

---

## 📁 `internal/plugins/git_dump.go` (COMPLETO)

```go
package plugins

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type GitDump struct{}

func (g *GitDump) Name() string {
	return "git_dump"
}

func (g *GitDump) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	base := strings.TrimRight(target.URL, "/")

	files := []string{
		"/.git/HEAD",
		"/.git/config",
		"/.git/index",
	}

	savePath := fmt.Sprintf("outputs/git/%s", sanitize(target.URL))
	os.MkdirAll(savePath, os.ModePerm)

	var downloaded int

	for _, f := range files {

		url := base + f

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", ctx.UserAgent)

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			continue
		}

		body, _ := io.ReadAll(resp.Body)

		if len(body) == 0 {
			continue
		}

		fileName := filepath.Join(savePath, filepath.Base(f))

		os.WriteFile(fileName, body, 0644)
		downloaded++
	}

	if downloaded > 0 {
		findings = append(findings, scanner.Finding{
			Title:       "Git Repository Dumped",
			Description: fmt.Sprintf("%d git files downloaded", downloaded),
			Severity:    "CRITICAL",
			Target:      target.URL,
			Score:       10.0,
			Confirmed:   true,
		})
	}

	return findings, nil
}

// 🔥 sanitiza nome da pasta
func sanitize(url string) string {
	url = strings.ReplaceAll(url, "https://", "")
	url = strings.ReplaceAll(url, "http://", "")
	url = strings.ReplaceAll(url, "/", "_")
	return url
}
````

---

# ⚠️ REGISTRAR O PLUGIN

## 📁 `internal/plugins/registry.go`

Adicione:

```go
r.Register(&GitDump{})
```

---

# 🚀 BUILD

```bash
go clean
go mod tidy
rm -f blueguard
CGO_ENABLED=0 go build -o blueguard ./cmd/blueguard
```

---

# ▶️ TESTE

```bash
./blueguard -u https://target-vulnerable.com
```

---

# 📂 RESULTADO

Se vulnerável:

```bash
outputs/git/example.com/
```

Com arquivos:

```bash
HEAD
config
index
```

---
