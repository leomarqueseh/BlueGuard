# 🛡️ BlueGuard

BlueGuard é uma ferramenta de segurança ofensiva focada em **detecção automatizada de vulnerabilidades web**, com arquitetura moderna, modular e preparada para evoluir para uma plataforma completa.

---

## 🚀 Features

- 🔍 Scanner de vulnerabilidades web
- ⚙️ Arquitetura modular (plugins)
- ⚡ Engine com concorrência (workers)
- 🌐 Dashboard web interativo
- 📊 Relatório HTML profissional
- 📦 Output em JSON (API-ready)
- 🌍 Suporte a múltiplos targets
- 🌐 Descoberta de subdomínios (base inicial)
- 🧠 Risk Engine com scoring

---

## 🧪 Vulnerabilidades detectadas

- Open Redirect
- Exposição de Headers
- Fingerprint de Tecnologias
- Git Exposure (em evolução)
- Subdomain Takeover (base)

---

## 🖥️ Dashboard Web

Interface web interativa para executar scans e visualizar resultados:

```bash
./blueguard -web
````

Acesse:

```
http://localhost:8080
```

### 🔥 Features do Dashboard

* Visual estilo Nessus
* Cards por vulnerabilidade
* Cores por severidade
* Contador de riscos (HIGH / MEDIUM / LOW / INFO)

---

## ⚙️ Uso via CLI

### Scan simples

```bash
./blueguard -u https://example.com
```

---

### Lista de targets

```bash
./blueguard -l targets.txt
```

---

### Descoberta de subdomínios

```bash
./blueguard -d example.com
```

---

### Output JSON

```bash
./blueguard -u https://example.com -json
```

---

### Relatório HTML

```bash
./blueguard -u https://example.com -html report.html
```

---

### Idioma

```bash
./blueguard -u https://example.com -lang pt-BR
```

---

## 🧠 Arquitetura

```
cmd/
 └── blueguard/

internal/
 ├── core/
 ├── scanner/
 ├── plugins/
 ├── worker/
 ├── risk/
 ├── report/
 ├── recon/
 ├── httpclient/
 └── dashboard/
```

---

## ⚡ Engine

* Execução paralela com workers
* Sistema baseado em plugins
* Pipeline:

```
Target → Plugins → Findings → Risk Engine → Output
```

---

## 📊 Exemplo de Output (JSON)

```json
[
  {
    "title": "Possible Open Redirect",
    "severity": "MEDIUM",
    "score": 6.5
  }
]
```

---

## 🛣️ Roadmap

### 🔥 Curto prazo

* [ ] Melhorar UI do dashboard
* [ ] Ordenação por severidade
* [ ] Expandir plugins

### 🚀 Médio prazo

* [ ] Banco de dados (SQLite)
* [ ] Histórico de scans
* [ ] API REST
* [ ] Melhorar subdomain discovery

### 🧠 Longo prazo

* [ ] Plataforma SaaS
* [ ] Multi-user
* [ ] Integração com ferramentas externas
* [ ] Machine learning para detecção

---

## 🧑‍💻 Desenvolvimento

```bash
git clone https://github.com/seu-user/BlueGuard
cd BlueGuard

go mod tidy
go build -o blueguard ./cmd/blueguard
```

---

## ⚠️ Aviso

Ferramenta destinada apenas para:

* Estudos
* Laboratórios
* Testes autorizados

Não utilize sem permissão.

---

## 📌 Status

🚧 Em desenvolvimento ativo

---

## ⭐ Contribuição

Pull requests são bem-vindos!

---

## 🏆 Autor

Desenvolvido por você 😈

````

---

# 🚀 AGORA FAZ ISSO

```bash
git add README.md
git commit -m "docs: complete README with dashboard, features and roadmap"
git push origin main
````

---

# 🔥 DICA (IMPORTANTE)

Cria uma pasta:

```bash
mkdir docs
```

Salva o print que você mandou como:

```bash
docs/dashboard.png
```

E adiciona no README:

```md
## 📊 Preview

![dashboard](./docs/dashboard.png)
