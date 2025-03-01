package web

import "github.com/gin-gonic/gin"

type GinWebEngine struct {
	*gin.Engine
}

func (g *GinWebEngine) NewGroup(prefix string) WebGroup {
	return &GinGroup{
		RouterGroup: g.Group(prefix),
	}
}

type GinGroup struct {
	*gin.RouterGroup
}

func (g *GinGroup) GET(path string, handler WebHandler) {
	g.RouterGroup.GET(path, func(c *gin.Context) {
		handler(&GinContext{c})
	})
}

// 实现POST/PUT/DELETE等其他方法...
