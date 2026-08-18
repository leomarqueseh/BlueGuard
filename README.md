Markdown
<h1 align="center">
  <br>
  🛡️ BlueGuard
  <br>
</h1>

<h4 align="center">Um scanner de vulnerabilidades modular e veloz, desenvolvido em Go.</h4>

<p align="center">
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Made with Go"></a>
  <a href="https://github.com/leomarqueseh/BlueGuard/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
  <a href="https://github.com/leomarqueseh/BlueGuard/releases"><img src="https://img.shields.io/github/v/release/leomarqueseh/BlueGuard" alt="Release"></a>
</p>

<p align="center">
  <a href="#-recursos">Recursos</a> •
  <a href="#-instalação">Instalação</a> •
  <a href="#-como-usar">Como Usar</a> •
  <a href="#-arquitetura-e-plugins">Arquitetura</a> •
  <a href="#-roadmap">Roadmap</a> •
  <a href="#-aviso-legal">Aviso Legal</a>
</p>

---

O **BlueGuard** é um scanner de vulnerabilidades inspirado em ferramentas do mercado como Nessus, mas projetado com o foco em velocidade, extensibilidade e usabilidade moderna. Ele combina um motor de reconhecimento inteligente com uma arquitetura baseada em plugins, executada de forma concorrente e controlada via CLI ou um Dashboard Web interativo.

## 🚀 Recursos

- **⚡ Alta Performance:** Escrito em Go com arquitetura de *Worker Pool* para execução assíncrona e concorrente.
- **🔌 Sistema de Plugins Extensível:** Adicione suas próprias regras de detecção facilmente.
- **🔍 Reconhecimento Profundo (Recon):** Mapeamento de subdomínios, resolução de IPs e scan de portas (via Nmap).
- **📊 Risk Engine Integrado:** Avaliação de score e severidade das vulnerabilidades.
- **🌐 Dashboard Web SaaS:** Interface intuitiva com mapas mentais interativos (estilo Maltego), seleção de plugins e painéis de risco.
- **📄 Múltiplos Outputs:** Exportação detalhada para CLI, JSON e HTML.

---

## ⚙️ Instalação

Certifique-se de ter o **[Go](https://golang.org/doc/install)** (versão 1.20 ou superior) instalado em seu ambiente.

```bash
# Clone o repositório
git clone [https://github.com/leomarqueseh/BlueGuard.git](https://github.com/leomarqueseh/BlueGuard.git)
cd BlueGuard

# Baixe as dependências
go mod tidy

# Compile o binário
go build -o blueguard ./cmd/blueguard
▶️ Como UsarO BlueGuard oferece controle total através de sua interface de linha de comando (CLI).Scans Básicos e AvançadosBash# Scan básico
./blueguard -u [https://example.com](https://example.com)

# Scan seguro (Evita execução de plugins destrutivos ou agressivos - RECOMENDADO)
./blueguard -u [http://testphp.vulnweb.com](http://testphp.vulnweb.com) -exclude git_dump,git_exposed

# Selecionar plugins específicos (Whitelist tem prioridade sobre Blacklist)
./blueguard -u [https://target.com](https://target.com) -plugins openredirect,techfingerprint

# Scan em lote a partir de um arquivo
./blueguard -l targets.txt

# Reconhecimento focado (Subdomínios + Expansão)
./blueguard -d example.com
Exportação de RelatóriosMesmo quando nenhuma vulnerabilidade crítica é encontrada, o BlueGuard gera relatórios sobre a superfície de ataque, serviços expostos e recomendações de hardening.Bash# Output em JSON (Ideal para integração com pipelines CI/CD)
./blueguard -u [https://example.com](https://example.com) -json

# Gerar relatório visual em HTML
./blueguard -u [https://example.com](https://example.com) -html report.html
🌐 Dashboard Web IntegradoInicie a interface web interativa para uma experiência visual completa:Bash./blueguard -web
Acesse http://localhost:8080 no seu navegador para visualizar:Painel de seleção e controle de plugins.Cards consolidados de severidade.Recon Map: Grafo interativo para exploração da superfície de ataque (expansão por clique, lazy loading).🧩 Arquitetura e PluginsO BlueGuard foi estruturado para ser altamente modular.Interface de PluginDesenvolver um novo plugin exige apenas a implementação de uma interface simples no Go:Gotype Plugin interface {
    Name() string
    Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error)
}
Plugins DisponíveisPluginDescriçãoopenredirectDetecta falhas de redirecionamento aberto.headerexposureIdentifica exposição de cabeçalhos sensíveis.techfingerprintMapeia tecnologias utilizadas (Servidores web, WAFs, CMS, etc).gitexposedVerifica a exposição indevida do diretório .git.git_dumpReconstrói e faz o dump de repositórios Git remotos expostos.Estrutura do ProjetoPlaintext.
├── cmd/
│   └── blueguard/main.go    # Entrypoint da aplicação
├── internal/
│   ├── core/                # Lógica central e orquestração
│   ├── scanner/             # Motor de varredura
│   ├── worker/              # Gerenciamento do Worker Pool (concorrência)
│   ├── plugins/             # Implementação dos plugins de detecção
│   ├── dashboard/           # Servidor web e UI
│   ├── httpclient/          # Cliente HTTP customizado
│   ├── recon/               # Módulos de reconhecimento e enumeração
│   ├── report/              # Geração de saídas (JSON, HTML)
│   └── risk/                # Engine de pontuação de risco
🧠 Roadmap[ ] Scan Profiles predefinidos (safe, full, aggressive).[ ] Crawler automático para mapeamento de rotas.[ ] Fuzzing inteligente de parâmetros.[ ] Fingerprint avançado para evasão de WAF e detecção de CMS.[ ] Persistência de dados (Integração com SQLite).[ ] Exportação completa de projetos em .ZIP.[ ] UI SaaS Premium (Melhorias no layout do dashboard).🤝 Como ContribuirPull requests são muito bem-vindos. Para mudanças importantes, abra uma issue primeiro para discutir o que você gostaria de alterar.Faça um Fork do projeto.Crie uma branch para sua feature: git checkout -b feature/nova-featureFaça o commit de suas alterações: git commit -m 'feat: Adicionando nova funcionalidade'Faça o Push para a branch: git push origin feature/nova-featureAbra um Pull Request.⚠️ Aviso LegalEste projeto foi desenvolvido exclusivamente para fins educacionais e testes de segurança autorizados. O autor não se responsabiliza pelo mau uso da ferramenta.🚫 NÃO utilize em alvos sem o consentimento explícito dos proprietários.🚫 NÃO execute plugins agressivos em ambientes de produção.2. Estratégia de Divulgação no LinkedInPara atrair a atenção de recrutadores e da comunidade de InfoSec, o ideal é focar nos desafios técnicos que você resolveu com o Go e na utilidade da ferramenta.Modelo de Postagem:🚀 Apresentando o BlueGuard: Um scanner de vulnerabilidades modular e de alta performance.Durante meus estudos em segurança ofensiva e desenvolvimento de ferramentas, percebi a necessidade de um scanner que unisse a velocidade das rotinas concorrentes do #Golang com uma interface intuitiva para análise visual de superfície de ataque.Foi assim que nasceu o BlueGuard. 🛡️Principais desafios e implementações técnicas:⚡ Worker Pools em Go: Arquitetura para garantir concorrência agressiva e alta performance sem estourar recursos de rede.🔌 Sistema de Plugins: Desenvolvi uma interface extensível (interface{}) onde qualquer pessoa pode criar módulos de detecção (ex: Open Redirect, Git Dump, Header Exposure).🌐 Dashboard & Grafo Inteligente: Não basta apenas gerar logs no terminal. Integrei um painel web que desenha a superfície de ataque em um mapa interativo (estilo Maltego).O projeto é 100% focado em testes autorizados, estudos de segurança e Bug Bounty.Gostaria do feedback da comunidade! Como vocês organizam a automação do fluxo de recon e scanning de vocês?🔗 Repositório no GitHub: [Insira o link aqui]#CyberSecurity #Golang #AppSec #RedTeam #BugBounty #InfoSec #OpenSource #DevSecOps
