
██████╗ ██╗     ██╗   ██╗███████╗ ██████╗ ██╗   ██╗ █████╗ ██████╗ ██████╗
██╔══██╗██║     ██║   ██║██╔════╝██╔════╝ ██║   ██║██╔══██╗██╔══██╗██╔══██╗
██████╔╝██║     ██║   ██║█████╗  ██║  ███╗██║   ██║███████║██████╔╝██║  ██║
██╔══██╗██║     ██║   ██║██╔══╝  ██║   ██║██║   ██║██╔══██║██╔══██╗██║  ██║
██████╔╝███████╗╚██████╔╝███████╗╚██████╔╝╚██████╔╝██║  ██║██║  ██║██████╔╝
╚═════╝ ╚══════╝ ╚═════╝ ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝

# 🛡️ BlueGuard – Automated Vulnerability Scanner

BlueGuard é um **framework modular de segurança ofensiva**, escrito em **Go**, focado em **alta confiabilidade**, **zero false positive** e **evolução contínua** para cobrir:

- Subdomain Takeover  
- Fuzzing inteligente  
- OWASP Top 10 (2021 + 2025)  

Inspirado em **Nuclei**, **Nessus** e **Burp Suite**, porém com **engine própria**, fingerprints em **YAML** e **controle total do fluxo de execução**.

---

## 🎯 Objetivos do Projeto

- Subdomain Takeover **100% confiável**
- Arquitetura extensível (**YAML + Go**)
- Execução rápida e silenciosa (**PASSIVE-LITE**)
- Evolução planejada para:
  - Fuzzing inteligente
  - OWASP Top 10
  - Dashboard web com gráficos, risco e severidade

---

## 🚀 Funcionalidades

### ✅ Implementado (Atual)

- **Subdomain Takeover Scanner**
- **HTTP Client avançado**
  - HTTPS first
  - Fallback HTTP
  - Redirect controlado
  - TLS Skip Verify (necessário para takeover)
- **Fingerprints em YAML**
  - CNAME validation
  - HTTP body match
- Timeout configurável
- Modo silencioso (**stealth**)
- Build estático (`CGO_ENABLED=0`)

---

### 🧪 Em Evolução / Planejado

- Engine **AND + negative**
- **Confidence score** (0–100)
- Output estruturado (**JSON / YAML**)
- Compatibilidade total com **Nuclei**
- **Fuzzing**:
  - Open Redirect
  - XSS
  - SQLi
  - IDOR
- Scripts customizados (pre / post / fuzz)
- **Dashboard web** (estilo Nessus)

---

## 📂 Estrutura do Projeto

```text
BlueGuard/
├── cmd/
│   └── blueguard/
│       └── main.go            # CLI principal
│
├── internal/
│   ├── analysis/
│   │   ├── takeover.go       # Engine takeover
│   │   └── loader.go         # Loader YAML
│   │
│   └── httpx/
│       ├── client.go         # HTTP client avançado
│       └── result.go         # Struct Result
│
├── fingerprints/
│   ├── aws_s3.yaml
│   ├── github_pages.yaml
│   ├── azure.yaml
│   └── heroku.yaml
│
├── outputs/                  # Gerado automaticamente
├── scripts/                  # (futuro – fuzz/custom)
├── web/                      # (futuro – dashboard)
├── go.mod
├── go.sum
└── README.md
````

---

## ⚙️ Instalação

### 1️⃣ Clonar o repositório

```bash
git clone https://github.com/leomarqueseh/BlueGuard.git
cd BlueGuard
```

### 2️⃣ Build (recomendado)

```bash
export CGO_ENABLED=0
go clean -cache
go mod tidy
go build -o blueguard ./cmd/blueguard
```

---

## 🏁 Uso Geral

```bash
./blueguard [flags]
```

---

## 🏳️ Flags Disponíveis

| Flag            | Descrição                           |
| --------------- | ----------------------------------- |
| `-t`            | Target único (ex: example.com)      |
| `-list`         | Arquivo `.txt` com domínios ou URLs |
| `-passive-lite` | Apenas takeover (rápido e seguro)   |
| `-passive`      | Recon passivo + takeover            |
| `-active`       | Scan ativo (em evolução)            |
| `-full`         | Scan completo (planejado)           |
| `-stealth`      | Modo silencioso                     |
| `-rate`         | Requests por segundo                |
| `-delay`        | Delay entre requests                |
| `-timeout`      | Timeout HTTP                        |

---

## 🟢 Modo PASSIVE-LITE (RECOMENDADO)

### 📌 O que faz?

* Apenas **Subdomain Takeover**
* Sem crawling
* Sem GAU
* Extremamente rápido
* Ideal para **bug bounty**

### ▶️ Usando lista

```bash
./blueguard -passive-lite -list subs.txt
```

**subs.txt**

```text
blog.example.com
cdn.example.com
static.example.com
```

### ▶️ Target único

```bash
./blueguard -passive-lite -t example.com
```

---

## 🟡 Modo PASSIVE

### 📌 O que faz?

* Reconhecimento passivo
* Coleta URLs
* Executa takeover

```bash
./blueguard -passive -t example.com
```

---

## 🔴 Modo FULL (Planejado)

```bash
./blueguard -full -t example.com
```

Inclui:

* Recon
* Takeover
* Fuzzing
* OWASP Top 10

---

## 🕵️ Modo STEALTH

Reduz logs e padrões agressivos.

```bash
./blueguard -passive-lite -list subs.txt -stealth
```

---

## 📤 Output Esperado

```text
🟢 Modo: PASSIVE-LITE (takeover only)
📄 Usando lista: subs.txt
✅ Nenhum takeover encontrado
```

Ou:

```text
🔥 POSSÍVEIS TAKEOVERS:
 - blog.example.com (AWS S3)
```

---

## 📄 Fingerprints YAML

📂 Local:

```text
fingerprints/
```

### Exemplo – `aws_s3.yaml`

```yaml
id: aws-s3
provider: AWS S3

cname:
  - amazonaws.com

http:
  matchers:
    - type: body
      words:
        - NoSuchBucket
        - The specified bucket does not exist
```

---

## 🔐 Confiabilidade

O BlueGuard **só reporta takeover** quando:

✔ CNAME válido
✔ Provider reconhecido
✔ Fingerprint HTTP específica
✔ HTTPS e HTTP testados
✔ Redirecionamento controlado

➡️ **Zero false positive por design**

---

## 🧠 Roadmap Oficial

### 🔵 Curto Prazo

* AND + negative engine
* Confidence score
* Output JSON / YAML
* Debug mode

### 🔴 Médio Prazo

* Fuzzing:

  * Open Redirect
  * XSS
  * SQLi
  * IDOR
* Scripts customizados
* OWASP Top 10 (2021 + 2025)

### 🟣 Longo Prazo

* Dashboard web (Nessus-like)
* Gráficos de risco
* Histórico de scans
* Classificação de severidade

---

## ⚠️ Aviso Legal

Este projeto deve ser utilizado **exclusivamente em ambientes autorizados**.
O autor não se responsabiliza por uso indevido.

---

## 📌 Status

🟢 **Ativo** | 🚧 **Em evolução** | 🧱 **Arquitetura sólida**

