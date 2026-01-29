Perfeito. Abaixo está um README.md COMPLETO, ATUALIZADO e PRONTO PARA O GITHUB, documentando tudo o que existe hoje no BlueGuard, com flags, exemplos reais, estrutura, boas práticas e status do projeto.

👉 Você pode copiar e colar diretamente como README.md na raiz do repositório.

⸻

🛡️ BlueGuard

BlueGuard é um framework de segurança ofensiva escrito em Go, focado inicialmente em Subdomain Takeover altamente confiável, com evolução planejada para OWASP Top 10, fuzzing inteligente e dashboard web estilo Nessus.

🎯 Objetivo: zero false positive, alta confiabilidade e extensibilidade via templates YAML e scripts customizados.

⸻

🚀 Funcionalidades atuais

✅ Implementado
	•	Subdomain Takeover Scanner
	•	Modo PASSIVE-LITE (rápido e seguro)
	•	Leitura de listas .txt
	•	HTTP Client avançado:
	•	HTTPS primeiro
	•	Fallback HTTP
	•	Redirects controlados
	•	Fingerprints em YAML
	•	Detecção por:
	•	CNAME
	•	Body HTTP
	•	Timeout configurável
	•	Compatível com Kali Linux
	•	Build sem CGO (CGO_ENABLED=0)

🧪 Em evolução
	•	Engine AND + negative
	•	Confidence score (0–100)
	•	Output JSON / YAML
	•	Compatibilidade total com Nuclei
	•	Fuzzing (open redirect, XSS, SQLi, etc.)
	•	Dashboard web com gráficos

⸻

📂 Estrutura do projeto

BlueGuard/
├── cmd/
│   └── blueguard/
│       └── main.go          # CLI principal
│
├── internal/
│   ├── analysis/
│   │   └── takeover.go     # Engine de takeover
│   │
│   └── httpx/
│       ├── client.go       # HTTP client (HTTPS + redirect)
│       └── result.go       # Struct Result
│
├── fingerprints/
│   ├── aws_s3.yaml
│   ├── github_pages.yaml
│   ├── azure.yaml
│   └── heroku.yaml
│
├── outputs/                # Gerado automaticamente
├── scripts/                # (planejado)
├── web/                    # (planejado)
├── go.mod
├── go.sum
└── README.md


⸻

⚙️ Instalação

1️⃣ Clonar o repositório

git clone https://github.com/leomarqueseh/BlueGuard.git
cd BlueGuard

2️⃣ Build (obrigatório sem CGO)

export CGO_ENABLED=0
go clean -cache
go mod tidy
go build -o blueguard ./cmd/blueguard


⸻

🏁 Uso básico

./blueguard [flags]


⸻

🟢 Modo PASSIVE-LITE (RECOMENDADO)

📌 O que faz?
	•	Apenas Subdomain Takeover
	•	Sem crawling
	•	Sem GAU
	•	Rápido
	•	Ideal para bug bounty e triagem inicial

⸻

▶️ Usar com lista de subdomínios

./blueguard -passive-lite -list subs.txt

📄 subs.txt (exemplo):

blog.example.com
cdn.example.com
static.example.com


⸻

▶️ Usar com target único

./blueguard -passive-lite -t example.com


⸻

📤 Output esperado

BlueGuard - Passive Takeover Scanner

🟢 Modo: PASSIVE-LITE (takeover only)
📄 Usando lista: subs.txt
✅ Nenhum takeover encontrado

Ou, se houver falha:

🔥 POSSÍVEIS TAKEOVERS:
 - blog.example.com (AWS S3)


⸻

🏳️ Flags disponíveis

Flag	Descrição
-t	Target único (ex: example.com)
-list	Arquivo .txt com domínios ou URLs
-passive-lite	Scan rápido (somente takeover)
-timeout	Timeout HTTP (default: 8s)


⸻

📄 Fingerprints (YAML)

📌 Local

fingerprints/

📌 Exemplo: aws_s3.yaml

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

🔐 Confiabilidade do Takeover

O BlueGuard considera takeover válido apenas quando:

✔ CNAME aponta para provider conhecido
✔ Resposta HTTP contém fingerprint específica
✔ HTTPS/HTTP testado
✔ Redirecionamentos controlados

❗ Sem brute force, sem guessing, sem falso positivo proposital.

⸻

🧠 Roadmap (próximos passos)

🔵 Curto prazo
	•	Engine AND + negative
	•	Confidence score real (0–100)
	•	Output JSON / YAML
	•	Debug mode

🔴 Médio prazo
	•	Fuzzing (open redirect, XSS, SQLi, IDOR)
	•	Scripts customizados
	•	OWASP Top 10 2021 + 2025

🟣 Longo prazo
	•	Dashboard web estilo Nessus
	•	Gráficos de risco
	•	Histórico de scans
	•	Classificação por ativo

⸻

⚠️ Aviso legal

Este projeto é destinado exclusivamente para fins educacionais e testes autorizados.
O uso indevido é de inteira responsabilidade do usuário.

⸻

🤝 Contribuições

Pull requests são bem-vindos.
Antes de contribuir:
	•	Siga o padrão YAML
	•	Evite false positives
	•	Documente fingerprints novos

⸻

📌 Status do projeto

🟢 Ativo e em evolução constante
🚀 Base sólida pronta para crescer

