package plugins

//
// 🔹 Registry central de plugins
//
type Registry struct {
	plugins []Plugin
}

//
// 🚀 Inicializa registry
//
func NewRegistry() *Registry {

	r := &Registry{}

	// 🔥 IMPORTANTE: usar ponteiro (&)
	r.Register(&OpenRedirect{})
	r.Register(&HeaderExposure{})
	r.Register(&TechFingerprint{})
	r.Register(&GitExposed{})
	r.Register(&GitDump{})

	return r
}

//
// 🔹 adiciona plugin
//
func (r *Registry) Register(p Plugin) {
	r.plugins = append(r.plugins, p)
}

//
// 🔥 retorna todos os plugins
//
func (r *Registry) All() []Plugin {
	return r.plugins
}
