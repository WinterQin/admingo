package orm

import "gorm.io/gorm"

type GORMAdapter struct {
	DB *gorm.DB
}

func (g *GORMAdapter) Create(model interface{}) error {
	return g.DB.Create(model).Error
}

func (g *GORMAdapter) FindByID(model interface{}, id interface{}) error {
	return g.DB.First(model, id).Error
}

// 实现其他接口方法...
