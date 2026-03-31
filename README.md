🛡️ BlueGuard

BlueGuard é um scanner de vulnerabilidades modular desenvolvido em Go, inspirado em ferramentas como o Nessus.

Projetado para ser rápido, extensível e moderno, o BlueGuard combina:

🔍 Reconhecimento de superfície (recon)
🔌 Arquitetura baseada em plugins
⚡ Execução concorrente
🌐 Dashboard web interativo estilo SaaS
🚀 Features
🔌 Sistema de plugins extensível
⚡ Worker pool (alta performance)
🌐 Dashboard web integrado
📊 Risk Engine (score + severidade)
🔍 Recon (subdomínios + IP + portas via Nmap)
🧠 Mapa mental interativo (estilo Maltego)
📄 Output:
CLI detalhado
JSON
HTML
🎯 Controle total de plugins via CLI
🧩 Plugins Disponíveis
Plugin	Descrição
openredirect	Detecta redirecionamentos abertos
headerexposure	Headers sensíveis expostos
techfingerprint	Tecnologias (Server, WAF, etc)
gitexposed	Verifica exposição do .git
git_dump	Reconstrói repositório Git remoto
⚙️ Instalação
git clone https://github.com/leomarqueseh/BlueGuard.git
cd BlueGuard

go mod tidy
go build -o blueguard ./cmd/blueguard
▶️ Uso
🔹 Scan básico
./blueguard -u https://example.com
🔹 Scan seguro (RECOMENDADO)
./blueguard -u http://testphp.vulnweb.com -exclude git_dump,git_exposed
🔹 Selecionar plugins específicos
./blueguard -u https://target.com -plugins openredirect,techfingerprint
🔹 Lista de targets
./blueguard -l targets.txt
🔹 Recon (subdomínios + expansão)
./blueguard -d example.com
🔹 Output JSON
./blueguard -u https://example.com -json
🔹 Relatório HTML
./blueguard -u https://example.com -html report.html
🌐 Dashboard Web
./blueguard -web

Acesse:

http://localhost:8080
✔ Interface inclui:
🎩 Logo hacker (chapéu + óculos)
🔌 Seleção de plugins
📊 Cards de severidade
📋 Tabela de vulnerabilidades
🧠 Recon Map interativo
🧠 Sistema de Plugins

Interface padrão:

type Plugin interface {
    Name() string
    Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error)
}

✔ Plugável
✔ Escalável
✔ Fácil de expandir

🎯 Controle de Plugins (CLI)
🔥 Incluir plugins (whitelist)
-plugins openredirect,techfingerprint
🔥 Excluir plugins (blacklist)
-exclude git_dump,git_exposed
🧠 Regra:
-plugins tem prioridade sobre -exclude
🧠 Recon Engine

Inspirado em ferramentas como:

Subfinder
Assetfinder
Nmap
🔥 Coleta:
Domínio
Subdomínios
IPs (DNS resolve)
Portas abertas
Serviços
Versões
🧠 Grafo Inteligente (Dashboard)

Estrutura:

example.com
 ├── subdomains
 │    ├── api.example.com
 │    ├── dev.example.com
 │
 ├── IP
 │    ├── 1.1.1.1
 │         ├── 80/http
 │         ├── 443/https
 │
 ├── Tech
 │    ├── nginx
 │    ├── PHP

✔ Expansão por clique
✔ Lazy loading (/expand)
✔ Estilo Maltego

📊 Relatórios (IMPORTANTE)

Mesmo quando nenhuma vulnerabilidade é encontrada, o BlueGuard gera:

✔ Informações coletadas
✔ Superfície de ataque
✔ Serviços expostos
✔ Recomendações de segurança

Exemplo:
Verificar headers de segurança
Ocultar versão do servidor
Restringir acesso a .git
Fechar portas desnecessárias
⚠️ Segurança & Uso Responsável

Este projeto é destinado para:

✔ Testes autorizados
✔ Estudos de segurança
✔ Laboratórios

🚫 NÃO use sem permissão
🚫 NÃO execute plugins agressivos em produção

🧠 Arquitetura
cmd/blueguard/main.go

internal/
  core/
  scanner/
  worker/
  plugins/
  dashboard/
  httpclient/
  recon/
  report/
  risk/
🧠 Roadmap
 Scan Profiles (safe / full / aggressive)
 Crawler automático
 Fuzzing de parâmetros
 Fingerprint avançado (WAF/CMS)
 Persistência (SQLite)
 Export ZIP
 Execução assíncrona
 UI SaaS nível premium
🤝 Contribuição

Pull requests são bem-vindos.

Fluxo recomendado:

git checkout -b feature/nova-feature
git commit -m "feat: descrição"
git push origin feature/nova-feature
📄 Licença

MIT License

⭐ Projeto em evolução

BlueGuard está evoluindo para se tornar uma plataforma completa combinando:

Scanner → Nessus
Recon visual → Maltego
Exploração → Burp Suite
