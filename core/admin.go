package core

type AdminConfig struct {
	ORM       ORMProvider
	WebEngine WebEngine
	Prefix    string
}

type Admin struct {
	config AdminConfig
	models []Model
}

func NewAdmin(config AdminConfig) *Admin {
	return &Admin{
		config: config,
		models: make([]Model, 0),
	}
}

func (a *Admin) Register(model Model) {
	a.models = append(a.models, model)
	a.registerRoutes(model)
}
