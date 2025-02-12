

# Admingo

Admingo 是一个轻量级的 Go 语言后台管理系统生成器，能够自动为您的数据模型生成 CRUD API。支持多种主流框架，目前已实现 Gin + GORM 的支持。

## 特性

- 🚀 自动 CRUD：自动生成增删改查接口
- 📝 模型自动发现：自动识别结构体字段
- 🛠 易于扩展：支持自定义配置
- 💡 零代码生成：无需手写 CRUD 代码

| 特性       | admingo | go-admin | gin-admin | gorm-admin |
| ---------- | ------- | -------- | --------- | ---------- |
| 轻量级     | ✅       | ❌        | ❌         | ✅          |
| 易用性     | ✅       | ❌        | ❌         | ✅          |
| UI界面     | 计划中  | ✅        | ✅         | ❌          |
| 多框架支持 | 计划中  | ✅        | ❌         | ❌          |
| 自动CRUD   | ✅       | ✅        | ✅         | ✅          |
| 配置复杂度 | 低      | 高       | 高        | 中         |
| 学习成本   | 低      | 高       | 高        | 中         |
| 可自定义性 | 高      | 中       | 中        | 低         |

## 安装

```bash
go get github.com/WinterQin/admingo
```

## 快速开始

### 1. 定义您的模型

```go
// Product 模型
type Product struct {
    gorm.Model     // ID  CreatedAt  UpdatedAt  DeletedAt
    Name        string
    Description string
    Price       float64
}

// User 模型
type User struct {
    gorm.Model
    Name        string
    Password    string
    Email       string
    IsAdmin     bool
}
```

### 2. 初始化并使用

```go
package main

import (
	"gorm.io/gorm"
    "github.com/gin-gonic/gin"
    "github.com/WinterQin/admingo"
)

func main() {
    // 初始化 Gin
    server := gin.Default()
    // 初始化 Gorm
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
    // 初始化 Admingo
    admin := admingo.NewAdmin(admingo.AdminConfig{
        DB:        db,        // 您的 GORM 实例
        URLPrefix: "/admin",  // API 前缀
        Engine:    server,    // Gin 实例
    })

    // 注册模型
    admin.Register(&Product{})
    admin.Register(&User{})

    // 启动服务器
    server.Run(":8888")
}
```

### 3. 生成的 API 端点

对于每个注册的模型，Admingo 会自动生成以下 RESTful API 端点：

```
POST   /admin/{模型名}     		- 创建新记录
GET    /admin/{模型名}     		- 获取记录列表
GET    /admin/{模型名}/:id 		- 获取单个记录
PUT    /admin/{模型名}/:id 		- 更新记录
DELETE /admin/{模型名}/:id 		- 删除记录
```

示例：
```
POST   /admin/Product           - 创建产品
GET    /admin/Product           - 获取产品列表
GET    /admin/Product/:id       - 获取单个产品
PUT    /admin/Product/:id       - 更新产品
DELETE /admin/Product/:id       - 删除产品
```

## 配置选项

```go
type AdminConfig struct {
    DB        *gorm.DB     // 数据库实例
    URLPrefix string       // API 路由前缀
    Engine    *gin.Engine  // Gin 实例
}
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 待实现功能

- [ ] 支持更多 Web 框架
- [ ] 支持更多数据库
- [ ] 添加权限控制
- [ ] 支持自定义验证规则
- [ ] 添加配套前端

