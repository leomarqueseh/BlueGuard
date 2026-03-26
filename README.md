# 🛡️ BlueGuard

BlueGuard é um scanner de vulnerabilidades modular desenvolvido em Go, inspirado em ferramentas profissionais como Nessus.

Projetado para ser rápido, extensível e fácil de evoluir, o BlueGuard utiliza arquitetura baseada em plugins e execução concorrente para análise de segurança em aplicações web.

---

## 🚀 Features

- 🔌 Arquitetura baseada em plugins
- ⚡ Execução concorrente (worker pool)
- 🌐 Dashboard web integrado (estilo SaaS)
- 📊 Risk Engine (classificação de severidade e score)
- 🔍 Recon (descoberta de subdomínios)
- 📄 Saída em JSON e HTML
- 🎯 Controle de plugins via CLI e Dashboard

---

## 🧩 Plugins Disponíveis

| Plugin            | Descrição |
|------------------|----------|
| openredirect     | Detecta possíveis redirecionamentos abertos |
| headerexposure   | Identifica headers sensíveis expostos |
| techfingerprint  | Detecta tecnologias (Server, WAF, etc) |
| gitexposed       | Verifica se `.git` está exposto |
| git_dump         | Reconstrói repositório Git remoto (avançado) |

---

## ⚙️ Instalação

```bash
git clone https://github.com/leomarqueseh/BlueGuard.git
cd BlueGuard

go mod tidy
go build -o blueguard ./cmd/blueguard
````

---

## ▶️ Uso

### 🔹 Scan básico

```bash
./blueguard -u https://example.com
```

---

### 🔹 Scan seguro (RECOMENDADO)

```bash
./blueguard -u http://testphp.vulnweb.com -exclude git_dump,gitexposed
```

---

### 🔹 Selecionar plugins

```bash
./blueguard -u https://target.com -plugins openredirect,techfingerprint
```

---

### 🔹 Lista de targets

```bash
./blueguard -l targets.txt
```

---

### 🔹 Descoberta de subdomínios

```bash
./blueguard -d example.com
```

---

### 🔹 Output JSON

```bash
./blueguard -u https://example.com -json
```

---

### 🔹 Relatório HTML

```bash
./blueguard -u https://example.com -html report.html
```

---

### 🌐 Dashboard Web

```bash
./blueguard -web
```

Acesse:

```
http://localhost:8080
```

✔ Interface estilo Nessus
✔ Seleção de plugins
✔ Visualização de resultados

---

## 🧠 Arquitetura

```
cmd/blueguard/main.go → EntryPoint

internal/
  core/        → contexto do scan
  scanner/     → engine de execução
  worker/      → concorrência
  plugins/     → sistema de plugins
  dashboard/   → interface web
  httpclient/  → cliente HTTP custom
  recon/       → descoberta
  report/      → geração de relatórios
  risk/        → análise de risco
```

---

## 🔥 Sistema de Plugins

Cada plugin segue a interface:

```go
type Plugin interface {
    Name() string
    Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error)
}
```

✔ Fácil de adicionar novos plugins
✔ Modular e escalável

---

## ⚠️ Segurança & Uso Responsável

Este projeto é destinado **exclusivamente para fins educacionais e testes autorizados**.

🚫 Não utilize em sistemas sem permissão
🚫 Não execute plugins agressivos em ambientes de produção

Ambiente recomendado para testes:

* [http://testphp.vulnweb.com](http://testphp.vulnweb.com)

---

## 🧠 Roadmap

* [ ] Scan Profiles (safe / full / aggressive)
* [ ] Crawler automático
* [ ] Fuzzing de parâmetros
* [ ] Fingerprint avançado (WAF, CMS, frameworks)
* [ ] Histórico de scans (SQLite)
* [ ] Exportação de relatórios (ZIP)
* [ ] UI SaaS premium

---

## 🤝 Contribuição

Pull requests são bem-vindos!

Para mudanças grandes:

1. Crie uma branch
2. Faça commits claros
3. Abra um PR

---

## 📄 Licença

MIT License

---

## ⭐ Projeto em evolução

BlueGuard está em constante evolução rumo a um scanner completo de segurança.


