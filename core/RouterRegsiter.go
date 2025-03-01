package core

import "path"

func (a *Admin) registerRoutes(model Model) {
	group := a.config.WebEngine.NewGroup(
		path.Join(a.config.Prefix, model.ModelName()),
	)

	// Create
	group.POST("/", a.createHandler(model))

	// Read (List)
	group.GET("/", a.listHandler(model))

	// Read (Single)
	group.GET("/:"+model.GetPrimaryKey(), a.getHandler(model))

	// Update
	group.PUT("/:"+model.GetPrimaryKey(), a.updateHandler(model))

	// Delete
	group.DELETE("/:"+model.GetPrimaryKey(), a.deleteHandler(model))
}
