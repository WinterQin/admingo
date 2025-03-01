package core

import "github.com/gin-gonic/gin"

func (a *Admin) createHandler(model Model) WebHandler {
	return func(c WebContext) {
		instance := model.NewInstance()
		if err := c.BindJSON(instance); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := a.config.ORM.Create(instance); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, instance)
	}
}
