package v010code

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"reflect"
)

type Admin struct {
	config     AdminConfig
	modelInfos map[string]*ModelInfo
	handler    *Handler
	engine     *gin.Engine
}

func NewAdmin(config AdminConfig) *Admin {
	return &Admin{
		config:     config,
		modelInfos: make(map[string]*ModelInfo),
		handler:    &Handler{db: config.DB},
		engine:     config.Engine,
	}
}

//func (a *Admin) getModelFields(model interface{}) []string {
//	// 获取结构体的类型
//	t := reflect.TypeOf(model).Elem()
//	// 创建一个切片来存储字段名称
//	var fields []string
//	// 递归获取字段
//	a.collectFields(t, &fields)
//	return fields
//}
//
//func (a *Admin) collectFields(t reflect.Type, fields *[]string) {
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		// 如果字段是嵌入的结构体，递归获取其字段
//		if field.Anonymous {
//			a.collectFields(field.Type, fields)
//		} else {
//			*fields = append(*fields, field.Name)
//		}
//	}
//}

func (a *Admin) getModelFields(model interface{}) []string {
	// 获取结构体的类型
	t := reflect.TypeOf(model).Elem()
	// 获取字段的数量
	numFields := t.NumField()
	// 创建一个切片来存储字段名称
	fields := make([]string, numFields)
	for i := 0; i < numFields; i++ {
		// 获取每个字段的名称
		fields[i] = t.Field(i).Name
	}
	return fields
}

func (a *Admin) registerRoutes(info *ModelInfo) {
	prefix := path.Join(a.config.URLPrefix, info.TableName)

	// 创建
	a.engine.POST(prefix, func(c *gin.Context) {
		a.handler.Create(c, info)
	})

	// 获取列表
	a.engine.GET(prefix, func(c *gin.Context) {
		a.handler.List(c, info)
	})

	// 获取单个
	a.engine.GET(prefix+"/:id", func(c *gin.Context) {
		a.handler.Get(c, info)
	})

	// 更新
	a.engine.PUT(prefix+"/:id", func(c *gin.Context) {
		a.handler.Update(c, info)
	})

	// 删除
	a.engine.DELETE(prefix+"/:id", func(c *gin.Context) {
		a.handler.Delete(c, info)
	})
}

func (a *Admin) Register(model interface{}) {
	// 使用反射获取模型信息
	t := reflect.TypeOf(model).Elem()

	modelInfo := &ModelInfo{
		Model:      model,
		TableName:  t.Name(),
		PrimaryKey: "id", // 默认主键
		Fields:     a.getModelFields(model),
	}
	fmt.Printf("Fields: %v\n", modelInfo.Fields)
	fmt.Printf("TableName: %v\n", modelInfo.TableName)

	// 存储模型信息
	a.modelInfos[modelInfo.TableName] = modelInfo

	//注册路由
	a.registerRoutes(modelInfo)
}
