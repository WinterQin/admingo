package v010code

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Model 接口定义了所有模型必须实现的方法
type Model interface {
	TableName() string
	GetID() uint
}

// Handler 处理所有通用CRUD操作
type Handler struct {
	db *gorm.DB
}

// AdminConfig 存储admin模块的配置
type AdminConfig struct {
	DB        *gorm.DB
	URLPrefix string // 默认 "/admingo"
	Engine    *gin.Engine
}

// ModelInfo 存储模型的元数据
type ModelInfo struct {
	Model      interface{}
	TableName  string
	PrimaryKey string
	Fields     []string
}
