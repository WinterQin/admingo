package itface

type WebContext interface {
	BindJSON(obj interface{}) error
	Param(key string) string
	JSON(code int, obj interface{})
}

type WebHandler func(WebContext)

type WebGroup interface {
	GET(path string, handler WebHandler)
	POST(path string, handler WebHandler)
	PUT(path string, handler WebHandler)
	DELETE(path string, handler WebHandler)
	Group(prefix string) WebGroup
}

type WebEngine interface {
	NewGroup(prefix string) WebGroup
}
