package itface

type Model interface {
	// 返回模型名称（用于路由生成）
	ModelName() string
	// 返回空模型实例（用于反射和创建新实例）
	NewInstance() interface{}
	// 返回主键字段名称（默认"ID"）
	GetPrimaryKey() string
}
