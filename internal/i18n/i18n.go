package i18n

var Lang = "en"

var messages = map[string]map[string]string{

	// 🔥 Open Redirect
	"open_redirect_title": {
		"en":    "Possible Open Redirect",
		"pt-BR": "Possível Redirecionamento Aberto",
	},
	"open_redirect_desc": {
		"en":    "Parameter may allow redirection",
		"pt-BR": "Parâmetro pode permitir redirecionamento",
	},

	// 🔥 Git Exposed
	"git_exposed_title": {
		"en":    "Git Repository Exposed",
		"pt-BR": "Repositório Git Exposto",
	},
	"git_exposed_desc": {
		"en":    ".git/config accessible",
		"pt-BR": ".git/config acessível",
	},

	// 🔥 Header Exposure
	"header_exposed_title": {
		"en":    "Server Header Exposed",
		"pt-BR": "Header do Servidor Exposto",
	},

	// 🔥 Fingerprint
	"fingerprint_title": {
		"en":    "Technology Fingerprint",
		"pt-BR": "Tecnologias Detectadas",
	},

	// 🔥 Severidade
	"severity_high": {
		"en":    "HIGH",
		"pt-BR": "ALTO",
	},
	"severity_medium": {
		"en":    "MEDIUM",
		"pt-BR": "MÉDIO",
	},
	"severity_low": {
		"en":    "LOW",
		"pt-BR": "BAIXO",
	},
	"severity_info": {
		"en":    "INFO",
		"pt-BR": "INFO",
	},

	// 🔥 Outros
	"detected": {
		"en":    "Detected",
		"pt-BR": "Detectado",
	},
}

func T(key string) string {

	if val, ok := messages[key]; ok {

		if msg, ok := val[Lang]; ok {
			return msg
		}

		if msg, ok := val["en"]; ok {
			return msg
		}
	}

	return key
}

// 🔥 Traduz severidade
func Severity(s string) string {

	switch s {
	case "HIGH":
		return T("severity_high")
	case "MEDIUM":
		return T("severity_medium")
	case "LOW":
		return T("severity_low")
	case "INFO":
		return T("severity_info")
	default:
		return s
	}
}
