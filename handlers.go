package admingo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

func (h *Handler) Create(c *gin.Context, info *ModelInfo) {
	// 创建模型的新实例
	instance := reflect.New(reflect.TypeOf(info.Model).Elem()).Interface()

	// 绑定请求数据
	if err := c.ShouldBindJSON(instance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存到数据库
	if err := h.db.Create(instance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, instance)
}

// 获取列表
func (h *Handler) List(c *gin.Context, info *ModelInfo) {
	var total int64
	// 创建模型的新实例
	modelType := reflect.TypeOf(info.Model).Elem()
	modelValue := reflect.New(modelType).Interface()

	// 创建结果切片
	resultsType := reflect.SliceOf(modelType)
	resultsValue := reflect.New(resultsType).Interface()

	// 分页参数
	currentPage := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	search := c.Query("search")
	// 转换为整数
	page, _ := strconv.Atoi(currentPage)
	size, _ := strconv.Atoi(pageSize)

	fmt.Println("page:", page, "pageSize:", size, "search:", search)

	// 构建查询
	query := h.db.Model(modelValue)
	query.Count(&total)

	// 如果有搜索关键词
	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	// 分页
	offset := (page - 1) * size
	query = query.Offset(offset).Limit(size)
	// 获取总数

	// 查询数据库
	if err := query.Find(resultsValue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将结果转换为正确的类型
	results := reflect.ValueOf(resultsValue).Elem().Interface()

	// 返回分页数据
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"records": results,
			"total":   total,
			"current": page,
			"size":    size,
		},
	})

}

func (h *Handler) Get(c *gin.Context, info *ModelInfo) {
	// 创建模型的新实例
	modelType := reflect.TypeOf(info.Model).Elem()
	modelValue := reflect.New(modelType).Interface()
	id := c.Param("id")

	if err := h.db.Where("id = ?", id).First(modelValue).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, modelValue)

}
func (h *Handler) Update(c *gin.Context, info *ModelInfo) {
	// 创建模型的新实例
	modelType := reflect.TypeOf(info.Model).Elem()
	modelValue := reflect.New(modelType).Interface()
	id := c.Param("id")

	if err := c.BindJSON(modelValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Model(modelValue).Where("id = ?", id).Updates(modelValue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
	fmt.Printf("test")
}
func (h *Handler) Delete(c *gin.Context, info *ModelInfo) {
	// 创建模型的新实例
	modelType := reflect.TypeOf(info.Model).Elem()
	modelValue := reflect.New(modelType).Interface()

	// 从 URL 中获取用户 ID
	id := c.Param("id")

	// 删除用户
	if err := h.db.Delete(modelValue, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回删除成功的消息
	c.JSON(http.StatusNoContent, nil)
}
