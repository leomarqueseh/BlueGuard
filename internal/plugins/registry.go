package plugins

import "strings"

//
// 🔹 Registry = container de plugins
//
type Registry struct {
	plugins []Plugin // 👈 usa a interface do plugin.go
}

//
// 🔹 Register adiciona plugin
//
func (r *Registry) Register(p Plugin) {
	r.plugins = append(r.plugins, p)
}

//
// 🔥 GetFiltered → filtra plugins (CLI + dashboard)
//
func (r *Registry) GetFiltered(include, exclude []string) []Plugin {

	var result []Plugin

	for _, p := range r.plugins {

		name := strings.ToLower(p.Name())

		// 🔹 include
		if len(include) > 0 {
			if !contains(include, name) {
				continue
			}
		}

		// 🔹 exclude
		if contains(exclude, name) {
			continue
		}

		result = append(result, p)
	}

	return result
}

//
// 🔹 helper
//
func contains(list []string, item string) bool {
	for _, v := range list {
		if strings.TrimSpace(v) == item {
			return true
		}
	}
	return false
}

//
// 🔥 Inicializa plugins
//
func NewRegistry() *Registry {

	r := &Registry{}

	// 🔐 DETECÇÃO
	r.Register(&OpenRedirect{})
	r.Register(&HeaderExposure{})
	r.Register(&TechFingerprint{}) // 👈 precisa existir

	// 🔥 EXPLOIT
	r.Register(&GitExposed{})
	r.Register(&GitDump{})

	return r
}
