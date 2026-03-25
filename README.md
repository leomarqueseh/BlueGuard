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
