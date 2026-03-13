package plugins

type Registry struct {
	plugins []Plugin
}

func NewRegistry() *Registry {

	r := &Registry{}

	r.Register(&GitExposed{})
	r.Register(&OpenRedirect{})
	r.Register(&HeaderExposure{})

	return r
}

func (r *Registry) Register(p Plugin) {
	r.plugins = append(r.plugins, p)
}

func (r *Registry) All() []Plugin {
	return r.plugins
}
