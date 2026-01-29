🛡️ BlueGuard – Automated Vulnerability Scanner

BlueGuard é um framework modular de segurança ofensiva, escrito em Go, focado em alta confiabilidade, zero false positive e evolução contínua para cobrir Subdomain Takeover, Fuzzing e OWASP Top 10 (2021 + 2025).

Inspirado em ferramentas como Nuclei, Nessus e Burp, porém com engine própria, fingerprints em YAML e controle total do fluxo.

⸻

🎯 Objetivos do Projeto
	•	Subdomain Takeover 100% confiável
	•	Arquitetura extensível (YAML + Go)
	•	Execução rápida (PASSIVE-LITE)
	•	Evolução para:
	•	Fuzzing inteligente
	•	OWASP Top 10
	•	Dashboard web com gráficos e severidade

⸻

🚀 Funcionalidades Implementadas

✅ Atuais
	•	Subdomain Takeover Scanner
	•	HTTP Client avançado:
	•	HTTPS first
	•	HTTP fallback
	•	Redirect controlado
	•	TLS skip verify (necessário para takeover)
	•	Fingerprints em YAML
	•	CNAME + HTTP Body match
	•	Timeout configurável
	•	Modo silencioso (stealth)
	•	Build estático (CGO_DISABLED)

🧪 Em evolução
	•	AND + negative engine
	•	Confidence score (0–100)
	•	Output JSON / YAML
	•	Compatibilidade total com Nuclei
	•	Fuzzing (open redirect, XSS, SQLi, IDOR…)
	•	Dashboard web (estilo Nessus)

⸻

📂 Estrutura do Projeto

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


⸻

⚙️ Instalação

1️⃣ Clonar o repositório

git clone https://github.com/leomarqueseh/BlueGuard.git
cd BlueGuard

2️⃣ Build (recomendado)

export CGO_ENABLED=0
go clean -cache
go mod tidy
go build -o blueguard ./cmd/blueguard


⸻

🏁 Uso Geral

./blueguard [flags]


⸻

🏳️ FLAGS DISPONÍVEIS

Flag	Descrição
-t	Target único (ex: example.com)
-list	Arquivo .txt com domínios ou URLs
-passive-lite	Somente takeover (rápido e seguro)
-passive	Recon passivo + takeover
-active	Scan ativo (em evolução)
-full	Scan completo (recon + active + fuzz – futuro)
-stealth	Modo silencioso (menos logs)
-rate	Requests por segundo
-delay	Delay entre requests
-timeout	Timeout HTTP


⸻

🟢 MODO PASSIVE-LITE (RECOMENDADO)

📌 O que faz?
	•	Apenas Subdomain Takeover
	•	Sem crawling
	•	Sem GAU
	•	Extremamente rápido
	•	Ideal para bug bounty

▶️ Usando lista

./blueguard -passive-lite -list subs.txt

📄 subs.txt

blog.example.com
cdn.example.com
static.example.com

▶️ Target único

./blueguard -passive-lite -t example.com


⸻

🟡 MODO PASSIVE

📌 O que faz?
	•	Recon passivo
	•	Coleta URLs
	•	Executa takeover

./blueguard -passive -t example.com


⸻

🔴 MODO FULL (planejado)

./blueguard -full -t example.com

Inclui:
	•	Recon
	•	Takeover
	•	Fuzzing
	•	OWASP Top 10

⸻

🕵️ MODO STEALTH

Reduz logs e padrões agressivos.

./blueguard -passive-lite -list subs.txt -stealth


⸻

📤 Output esperado

🟢 Modo: PASSIVE-LITE (takeover only)
📄 Usando lista: subs.txt
✅ Nenhum takeover encontrado

Ou:

🔥 POSSÍVEIS TAKEOVERS:
 - blog.example.com (AWS S3)


⸻

📄 Fingerprints YAML

📂 Local:

fingerprints/

Exemplo – aws_s3.yaml

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


⸻

🔐 Confiabilidade

O BlueGuard só reporta takeover quando:

✔ CNAME válido
✔ Provider reconhecido
✔ HTTP fingerprint específica
✔ HTTPS/HTTP testados
✔ Redirecionamento controlado

➡️ Zero false positive por design

⸻

🧠 Roadmap Oficial

🔵 Curto prazo
	•	AND + negative engine
	•	Confidence score
	•	Output JSON / YAML
	•	Debug mode

🔴 Médio prazo
	•	Fuzzing:
	•	Open Redirect
	•	XSS
	•	SQLi
	•	IDOR
	•	Scripts customizados
	•	OWASP Top 10 (2021 + 2025)

🟣 Longo prazo
	•	Dashboard web (Nessus-like)
	•	Gráficos de risco
	•	Histórico de scans
	•	Classificação de severidade

⸻

⚠️ Aviso Legal

Uso exclusivamente autorizado.
O autor não se responsabiliza por uso indevido.

⸻

📌 Status

🟢 Ativo | Em evolução | Arquitetura sólida

